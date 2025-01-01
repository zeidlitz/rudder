package server

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/zeidlitz/rudder/loadbalancer"
)

type Server struct {
	serverlist []string
	lb         loadbalancer.LoadBalancer
}

func (s *Server) loadBalanceHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, clientRequest *http.Request) {
		server, err := s.lb.GetServer()
		if err != nil {
			slog.Error("Internal error", "message", err.Error())
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		slog.Info("Forwarding to", "server", server)
		relay(w, clientRequest, server)
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

func relay(w http.ResponseWriter, clientRequest *http.Request, server string) {
	url := server + clientRequest.RequestURI
	serverRequest, err := http.NewRequest(clientRequest.Method, url, clientRequest.Body)
	if err != nil {
		slog.Error("Error when creating new request", "error", err.Error())
		return
	}

	for key, values := range clientRequest.Header {
		for _, value := range values {
			serverRequest.Header.Add(key, value)
		}
	}

	slog.Info("Sending ", "request", serverRequest)
	client := &http.Client{}
	resp, err := client.Do(serverRequest)

	if err != nil {
		slog.Error("Error making request to the destination server", "error", err.Error())
		http.Error(w, "Error making request to the destination server", http.StatusBadGateway)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Error reading response body: ", "err", err)
		http.Error(w, "Error reading resposne body", http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		slog.Error("Error unmarshalling response body JSON: ", "err", err)
		http.Error(w, "Error unrmarshalling resposne body JSON", http.StatusInternalServerError)
		return
	}

	slog.Info("Response from ", "server", server, "response", data)
	relayResponse, err := json.Marshal(data)
	if err != nil {
		slog.Error("Error marshalling response body JSON: ", "err", err)
		http.Error(w, "Error marshalling resposne body JSON", http.StatusInternalServerError)
		return
	}
	w.Write(relayResponse)

	defer resp.Body.Close()
}

func (s *Server) Start(host string) error {
	http.HandleFunc("/", s.loadBalanceHandler())
	http.ListenAndServe(host, nil)
	return nil
}
