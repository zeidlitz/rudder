package server

import (
	"log/slog"
	"net/http"
	"strings"
  "encoding/json"

	"github.com/zeidlitz/rudder/loadbalancer"
)

type Server struct {
	serverlist []string
	lb         loadbalancer.LoadBalancer
}

func (s *Server) loadBalanceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chosenServer, err := s.lb.GetServer()
		if err != nil {
			slog.Error("Internal error", "message", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
    slog.Info("Forwarding", "server", chosenServer)
    relay(w,r,chosenServer)
	}
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

func relay(w http.ResponseWriter, r *http.Request, destination string) {
  url := destination + r.RequestURI
  req, err := http.NewRequest(r.Method, url, r.Body)
  if err != nil {
    slog.Error("Error when creating new request", "error", err.Error())
    return
  }

  // Copy headers from original request to the new request
  for key, values := range r.Header {
    for _, value := range values {
      req.Header.Add(key, value)
    }
  }

  slog.Info("Sending request", "request", req)
  client := &http.Client{}
  resp, err := client.Do(req)

  if err != nil {
    slog.Error("Error making request to the destination server", "error", err.Error())
    http.Error(w, "Error making request to the destination server", http.StatusBadGateway)
    return
  }
  slog.Info("Response from server", "status code", resp.StatusCode)
  // TODO: Make the response me actual response from the request, add err handling
  response := map[string]string{"server": "running", "status": "200 OK"}
	jsonResponse, err := json.Marshal(response)
  w.Write(jsonResponse)
  defer resp.Body.Close()
}

func (s *Server) Start(host string) error {
	http.HandleFunc("/", s.loadBalanceHandler())
	http.ListenAndServe(host, nil)
	return nil
}
