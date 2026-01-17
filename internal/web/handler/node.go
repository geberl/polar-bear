package handler

import (
	"net/http"
	"net/url"
	"time"

	"polar-bear/internal/config"
	"polar-bear/internal/core"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/store"
	"polar-bear/internal/web/view/node"
)

func Node(
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

			no, err := url.QueryUnescape(r.PathValue("no"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			}

			res := core.GetNode(store, no)
			nss := core.GetNamespaces(store)

			err = node.DetailView(&startTime, cfg, rm, no, res, nss).Render(r.Context(), w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}
