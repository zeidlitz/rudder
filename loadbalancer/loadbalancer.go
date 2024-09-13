package loadbalancer

import "fmt"

type LoadBalancer interface {
	Configure(servers []string) error
	GetServer() (string, error)
}

func GetLoadBalancer(algorithm string, serverlist []string) (LoadBalancer, error) {
	switch algorithm {
	case "roundrobin":
		return &RoundRobin{currentIndex: 0, servers: serverlist}, nil
	case "iphash":
		return &IPHash{servers: serverlist}, nil
	default:
		return nil, fmt.Errorf("unknown load balancing algorithm: %s", algorithm)
	}
}
