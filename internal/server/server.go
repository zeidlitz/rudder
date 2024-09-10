package server

import (
	"log/slog"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handeled request", "host", r.Host)
	w.WriteHeader(http.StatusOK)
}

func Start(host string) error {
	slog.Info("Starting up", "host", host)
	http.HandleFunc("/", handler)
	http.ListenAndServe(host, nil)
	return nil
}
