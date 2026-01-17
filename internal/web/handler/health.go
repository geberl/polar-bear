package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"polar-bear/internal/runtimemeta"
)

func Health(
	rm *runtimemeta.RuntimeMeta,
) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			uptime := time.Since(rm.StartTime)

			health := struct {
				Status string `json:"status"`
				Uptime string `json:"uptime"`
				Time   string `json:"time"`
			}{
				Status: "ok",
				Uptime: uptime.String(),
				Time:   time.Now().UTC().Format(time.RFC3339),
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			_ = json.NewEncoder(w).Encode(health)
		},
	)
}
