package server

import (
	"compress/flate"
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

type CompressionType string

const (
	BR      CompressionType = "br"
	GZIP    CompressionType = "gzip"
	DEFLATE CompressionType = "deflate"
)

type compressedResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *compressedResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func getCompressionType(acceptEncoding string) CompressionType {
	for encoding := range strings.SplitSeq(acceptEncoding, ",") {
		trimmed := strings.TrimSpace(encoding)
		switch trimmed {
		case string(BR):
			return BR
		case string(GZIP):
			return GZIP
		case string(DEFLATE):
			return DEFLATE
		}
	}
	return ""
}

func compressionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch getCompressionType(r.Header.Get("Accept-Encoding")) {
		case BR:
			bw := brotli.NewWriterLevel(w, brotli.DefaultCompression)
			defer bw.Close()

			w.Header().Set("Content-Encoding", string(BR))
			w.Header().Set("Vary", "Accept-Encoding")

			cw := &compressedResponseWriter{
				Writer:         bw,
				ResponseWriter: w,
			}
			next.ServeHTTP(cw, r)
		case GZIP:
			gz := gzip.NewWriter(w)
			defer gz.Close()

			w.Header().Set("Content-Encoding", string(GZIP))
			w.Header().Set("Vary", "Accept-Encoding")

			cw := &compressedResponseWriter{
				Writer:         gz,
				ResponseWriter: w,
			}
			next.ServeHTTP(cw, r)
		case DEFLATE:
			fw, err := flate.NewWriter(w, flate.DefaultCompression)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			defer fw.Close()

			w.Header().Set("Content-Encoding", string(DEFLATE))
			w.Header().Set("Vary", "Accept-Encoding")

			cw := &compressedResponseWriter{
				Writer:         fw,
				ResponseWriter: w,
			}
			next.ServeHTTP(cw, r)
		default:
			next.ServeHTTP(w, r)
		}
	})
}
