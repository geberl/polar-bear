package server

import (
	"log/slog"
	"net/http"
	"strings"
	"time"
)

type responseRecorder struct {
	http.ResponseWriter
	status   int
	numBytes int
}

func (re *responseRecorder) WriteHeader(code int) {
	re.status = code
	re.ResponseWriter.WriteHeader(code)
}

func (re *responseRecorder) Write(data []byte) (int, error) {
	re.numBytes += len(data)
	return re.ResponseWriter.Write(data)
}

func loggingMiddleware(next http.Handler) http.Handler {
	logger := slog.With("component", "request-logger")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := responseRecorder{ResponseWriter: w, status: http.StatusOK}
		r.Header.Set("X-Request-Time", start.Format(time.RFC3339Nano))
		next.ServeHTTP(&rw, r)

		headers := map[string]string{}
		for name, values := range r.Header {
			headers[name] = strings.Join(values, ", ")
		}
		logger.DebugContext(r.Context(),
			"incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"query", r.URL.RawQuery,
			"ip", r.RemoteAddr,
			"bytes", rw.numBytes,
			"status", rw.status,
			"latency", time.Since(start).String(),
			"headers", headers,
		)
	})
}
