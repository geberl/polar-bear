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
	"polar-bear/internal/web/view/replicaset"
	"polar-bear/internal/web/view/statefulset"
)

func Resources(
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

			nss := core.GetNamespaces(store)

			switch res {
			case "pd":
				pds := core.GetPods(store, ns)
				err = pod.ListView(&startTime, cfg, rm, ns, pds, nss).Render(r.Context(), w)
			case "rs":
				rss := core.GetReplicaSets(store, ns)
				err = replicaset.ListView(&startTime, cfg, rm, ns, rss, nss).Render(r.Context(), w)
			case "sts":
				sts := core.GetStatefulSets(store, ns)
				err = statefulset.ListView(&startTime, cfg, rm, ns, sts, nss).Render(r.Context(), w)
			case "deploy":
				deploys := core.GetDeployments(store, ns)
				err = deployment.ListView(&startTime, cfg, rm, ns, deploys, nss).Render(r.Context(), w)
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}
