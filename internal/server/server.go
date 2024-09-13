package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/zeidlitz/go-balancer/loadbalancer"
)

type Server struct {
	serverlist []string
	lb         loadbalancer.LoadBalancer
}

func (s *Server) loadBalanceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chosenServer, err := s.lb.GetServer()
		if err != nil {
			slog.Error("Error", "error", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		slog.Info("Forwarding to server", "server", chosenServer)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Handeled request", "host", r.Host)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) Configure(algorithm string, serverlist string) error {
	servers := strings.Split(serverlist, ", ")
	lb, err := loadbalancer.GetLoadBalancer(algorithm, servers)
	if err != nil {
		slog.Error("Error when configuring server", "error", err.Error())
		return err
	}

	s.lb = lb
	return nil
}

func (s *Server) Start(host string) error {
	slog.Info("Starting up", "host", host)
	http.HandleFunc("/", s.loadBalanceHandler())
	http.ListenAndServe(host, nil)
	return nil
}
