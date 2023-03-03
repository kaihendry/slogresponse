package slogresponse

import (
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

func New(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			slog.Info("slogresponse",
				"req_method", r.Method,
				"req_ip", r.RemoteAddr,
				"req_path", r.RequestURI,
				slog.Duration("duration", time.Since(start)),
			)
		}()
		h.ServeHTTP(w, r)
	})
}
