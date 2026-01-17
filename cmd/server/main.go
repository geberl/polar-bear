package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/peterbourgon/ff"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"

	"polar-bear/cmd"
	"polar-bear/internal/config"
	"polar-bear/internal/event"
	"polar-bear/internal/informer"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/server"
	"polar-bear/internal/store"
)

const (
	applicationName = "polar-bear"
	envVarPrefix    = "POLAR_BEAR"
)

// Version is set at build time using ldflags.
var Version string

func main() {
	fs := flag.NewFlagSet(applicationName, flag.ExitOnError)
	ll := fs.String("loglevel", "info", "Verbosity of logging, one of debug/info/warn/error")
	lf := fs.String("logformat", "human", "Format of logging, one of human/json")
	cn := fs.String("cluster-name", "My Cluster", "name of the cluster")
	dm := fs.Bool("devmode", false, "Use non-optimized Tailwind CSS file with all classes")
	hn := fs.String("hostname", "My Host", "name of the host that serves the application")
	hl := fs.String("http-listen-address", "localhost:8888", "http listen address")
	ml := fs.String("metrics-listen-address", "localhost:8889", "metrics listen address")
	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVarPrefix(envVarPrefix))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger, err := cmd.MakeLogger(*ll, *lf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	slog.SetDefault(logger)
	slog.Info(
		"starting",
		"name", applicationName,
		"loglevel", *ll,
		"logformat", *lf,
	)

	rm, err := runtimemeta.GetRuntimeMeta(Version, *hn)
	if err != nil {
		slog.Error(
			"unable to get runtime metadata",
			"error", err,
		)
		os.Exit(1)
	}
	slog.Info(
		"runtime metadata",
		"hostname", rm.HostName,
		"path", rm.Path,
		"start_time", rm.StartTime,
		"version", rm.Version,
		"revision", rm.Revision,
		"revision_short", rm.RevisionShort,
		"last_commit", rm.LastCommit,
		"dirty_build", rm.DirtyBuild,
		"go_version", rm.GoVersion,
		"go_arch", rm.GoArch,
		"go_os", rm.GoOS,
	)

	cfg := &config.Config{
		ClusterName:          *cn,
		DevMode:              *dm,
		HTTPListenAddress:    *hl,
		MetricsListenAddress: *ml,
	}
	slog.Info(
		"config",
		"cluster_name", cfg.ClusterName,
		"dev_mode", cfg.DevMode,
		"http_listen_address", cfg.HTTPListenAddress,
		"metrics_listen_address", cfg.MetricsListenAddress,
	)

	ctx := context.Background()
	if err := run(ctx, rm, cfg); err != nil {
		slog.Error("failed to start server", "err", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	rm *runtimemeta.RuntimeMeta,
	cfg *config.Config,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	store, err := store.NewOtterStore()
	if err != nil {
		return fmt.Errorf("failed to create new store: %v", err)
	}

	ed, err := event.NewDistributer(slog.With("component", "event-distributer"))
	if err != nil {
		return fmt.Errorf("failed to create new event distributer: %v", err)
	}

	fct, err := informer.NewInformerFactory("")
	if err != nil {
		return fmt.Errorf("failed to create new informer factory: %v", err)
	}

	infs := []informer.Informer{
		informer.NewNodeInformer(fct, store, ed),
		informer.NewNamespaceInformer(fct, store, ed),
		informer.NewPodInformer(fct, store, ed),
		informer.NewReplicaSetInformer(fct, store, ed),
		informer.NewStatefulSetInformer(fct, store, ed),
		informer.NewDeploymentInformer(fct, store, ed),
		informer.NewServiceInformer(fct, store, ed),
		informer.NewIngressInformer(fct, store, ed),
	}

	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	srv := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		IdleTimeout:       120 * time.Second,
		Addr:              cfg.HTTPListenAddress,
		Handler:           server.GetRoutes(rm, cfg, mdlw, store, ed),
	}

	metricsSrv := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		IdleTimeout:       60 * time.Second,
		Addr:              cfg.MetricsListenAddress,
		Handler:           promhttp.Handler(),
	}

	infCount := len(infs)
	errChan := make(chan error, infCount+2) // http server, metrics server

	for _, inf := range infs {
		go func() {
			slog.InfoContext(ctx, "informer running", "kind", inf.Kind())
			if err := inf.Run(); err != nil && err != http.ErrServerClosed {
				slog.ErrorContext(ctx, "failed to start informer",
					"kind", inf.Kind(),
					"err", err,
				)
				errChan <- err
			}
		}()
	}

	go func() {
		slog.InfoContext(ctx, "http server running", "address", cfg.HTTPListenAddress)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "failed to start http server", "err", err)
			errChan <- err
		}
	}()

	go func() {
		slog.InfoContext(ctx, "metrics server running", "address", cfg.MetricsListenAddress)
		if err := metricsSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.ErrorContext(ctx, "failed to start metrics server", "err", err)
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		slog.InfoContext(ctx, "shutting down servers")
	}

	var errs []error

	shutdownCtx := context.Background()
	shutdownCtx, shutdownCancel := context.WithTimeout(shutdownCtx, 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.ErrorContext(shutdownCtx, "failed to shutdown http server gracefully", "err", err)
		errs = append(errs, err)
	}
	slog.InfoContext(ctx, "http server shut down successfully", "address", cfg.HTTPListenAddress)

	if err := metricsSrv.Shutdown(shutdownCtx); err != nil {
		slog.ErrorContext(shutdownCtx, "failed to shutdown metrics server gracefully", "err", err)
		errs = append(errs, err)
	}
	slog.InfoContext(ctx, "metrics server shut down successfully", "address", cfg.MetricsListenAddress)

	for _, inf := range infs {
		if err := inf.Close(); err != nil {
			slog.ErrorContext(shutdownCtx, "failed to shutdown informer gracefully",
				"kind", inf.Kind(),
				"err", err,
			)
			errs = append(errs, err)
		}
		slog.InfoContext(ctx, "informer shut down successfully", "kind", inf.Kind())
	}

	if err := store.Close(); err != nil {
		slog.ErrorContext(shutdownCtx, "failed to shutdown store gracefully", "err", err)
		errs = append(errs, err)
	}
	slog.InfoContext(ctx, "db shut down successfully")

	return errors.Join(errs...)
}
