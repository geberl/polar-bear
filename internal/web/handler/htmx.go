package handler

import (
	"net/http"

	"polar-bear/internal/core"
	"polar-bear/internal/runtimemeta"
	"polar-bear/internal/store"
	"polar-bear/internal/web/view/shared"
)

func HTMXOpenSidebar(
	rm *runtimemeta.RuntimeMeta,
	store store.Store,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			nss := core.GetNamespaces(store)

			err := shared.SidebarState("open", rm, nss, "", "").Render(r.Context(), w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}

func HTMXCloseSidebar(
	rm *runtimemeta.RuntimeMeta,
	store store.Store,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			nss := core.GetNamespaces(store)

			err := shared.SidebarState("closed", rm, nss, "", "").Render(r.Context(), w)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		},
	)
}
