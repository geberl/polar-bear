package server

import (
	"net/http"

	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"

	"polar-bear/internal/config"
	"polar-bear/internal/event"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/store"
	"polar-bear/internal/web/handler"
)

func GetRoutes(
	rm *runtimemeta.RuntimeMeta,
	cfg *config.Config,
	metricsMiddleware middleware.Middleware,
	store store.Store,
	event event.Distribution,
) http.Handler {
	// Routes registered in this mux WILL include middlewares
	mwMux := http.NewServeMux()

	mwMux.Handle("GET /no", handler.Nodes(cfg, rm, store))
	mwMux.Handle("GET /no/", handler.Nodes(cfg, rm, store))
	mwMux.Handle("GET /no/{no}", handler.Node(cfg, rm, store))
	mwMux.Handle("GET /no/{no}/", handler.Node(cfg, rm, store))

	mwMux.Handle("GET /ns/{ns}", handler.Namespace(cfg, rm, store))
	mwMux.Handle("GET /ns/{ns}/", handler.Namespace(cfg, rm, store))

	mwMux.Handle("GET /ns/{ns}/{res}", handler.Resources(cfg, rm, store))
	mwMux.Handle("GET /ns/{ns}/{res}/", handler.Resources(cfg, rm, store))
	mwMux.Handle("GET /ns/{ns}/{res}/{name}", handler.Resource(cfg, rm, store))
	mwMux.Handle("GET /ns/{ns}/{res}/{name}/", handler.Resource(cfg, rm, store))

	mwMux.Handle("GET /health", handler.Health(rm))
	mwMux.Handle("GET /info", handler.Info(cfg, rm, store))

	mwMux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mwMux.Handle("GET /_open-sidebar", handler.HTMXOpenSidebar(rm, store))
	mwMux.Handle("GET /_close-sidebar", handler.HTMXCloseSidebar(rm, store))

	mwMux.Handle("GET /", handler.Cluster(cfg, rm, store))

	// Middlewares to apply
	mwHnd := loggingMiddleware(mwMux)
	mwHnd = compressionMiddleware(mwHnd)
	mwHnd = std.Handler("", metricsMiddleware, mwHnd)

	// Routes registered directly in this mux WILL NOT include middlewares
	rootMux := http.NewServeMux()
	rootMux.Handle("/", mwHnd)

	rootMux.Handle("GET /ws/{ns}", handler.Websocket(event, store))

	return rootMux
}
