package slogresponse

import (
	"net/http"

	"golang.org/x/exp/slog"
)

func New(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("slogresponse",
			"req_method", r.Method,
			"req_ip", r.RemoteAddr,
			"req_path", r.RequestURI,
		)
		h.ServeHTTP(w, r)
	})
}
