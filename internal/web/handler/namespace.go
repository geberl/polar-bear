package handler

import (
	"net/http"
	"net/url"
	"time"

	"polar-bear/internal/config"
	"polar-bear/internal/core"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/store"
	"polar-bear/internal/web/view/namespace"
)

func Namespace(
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

			data := &namespace.Data{
				Start:            &startTime,
				Config:           cfg,
				Meta:             rm,
				Namespace:        core.GetNamespace(store, ns),
				Namespaces:       core.GetNamespaces(store),
				PodCount:         core.CountPods(store, ns),
				ReplicaSetCount:  core.CountReplicaSets(store, ns),
				StatefulSetCount: core.CountStatefulSets(store, ns),
				DeploymentCount:  core.CountDeployments(store, ns),
				ServiceCount:     core.CountServices(store, ns),
				IngressCount:     core.CountIngresses(store, ns),
			}

			err = namespace.DetailView(data).Render(r.Context(), w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}
