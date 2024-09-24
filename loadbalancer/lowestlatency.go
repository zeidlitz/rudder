package loadbalancer

import (
	"log/slog"
	"net"
	"time"
)


type LowestLatency struct {
	servers []string
}

func getLatency(host string) (time.Duration, error) {
	start := time.Now()
	conn, err := net.DialTimeout("tcp", host+":80", time.Millisecond)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	latency := time.Since(start)
	return latency, nil
}

func (ll *LowestLatency) Configure(serverlist []string) error {
	ll.servers = serverlist
	return nil
}

func (ll *LowestLatency) GetServer() (string, error) {
  var lowestLatencyServer string
  var lowestLatency time.Duration
  var err error
  
  for _, server := range ll.servers {
  var latency time.Duration
  latency, err = getLatency(server)
  if err != nil {
      slog.Warn("Failed to reach host", server, err )
      continue
    }
    slog.Info("Latency to %s: %v", server, latency)
    if lowestLatencyServer == "" || latency < lowestLatency {
      lowestLatencyServer = server
      lowestLatency = latency
    }
  }
  return lowestLatencyServer, err
}
