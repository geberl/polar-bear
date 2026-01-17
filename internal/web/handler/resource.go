package handler

import (
	"net/http"
	"net/url"
	"time"

	"polar-bear/internal/config"
	"polar-bear/internal/core"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/store"
	"polar-bear/internal/web/view/deployment"
	"polar-bear/internal/web/view/pod"
)

func Resource(
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

			ns, err := url.QueryUnescape(r.PathValue("ns"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			}

			res, err := url.QueryUnescape(r.PathValue("res"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			}

			name, err := url.QueryUnescape(r.PathValue("name"))
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotAcceptable)
				return
			}

			nss := core.GetNamespaces(store)

			switch res {
			case "pd":
				pd := core.GetPod(store, ns, name)
				err = pod.DetailView(&startTime, cfg, rm, ns, name, pd, nss).Render(r.Context(), w)
			case "deploy":
				deploy := core.GetDeployment(store, ns, name)
				err = deployment.DetailView(&startTime, cfg, rm, ns, name, deploy, nss).Render(r.Context(), w)
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}
