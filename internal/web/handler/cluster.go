package handler

import (
	"net/http"
	"time"

	"polar-bear/internal/config"
	"polar-bear/internal/core"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/store"
	"polar-bear/internal/web/view/cluster"
)

func Cluster(
	cfg *config.Config,
	rm *runtimemeta.RuntimeMeta,
	store store.Store,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := r.Header.Get("X-Request-Time")
			startTime, err := time.Parse(time.RFC3339Nano, start)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			nss := core.GetNamespaces(store)

			err = cluster.View(&startTime, cfg, rm, nss).Render(r.Context(), w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}
