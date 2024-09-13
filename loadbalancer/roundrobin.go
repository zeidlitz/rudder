package loadbalancer

type RoundRobin struct {
	servers      []string
	currentIndex int
}

func (rr *RoundRobin) Configure(serverlist []string) error {
	rr.servers = serverlist
	return nil
}

func (rr *RoundRobin) GetServer() (string, error) {
	server := rr.servers[rr.currentIndex]
	rr.currentIndex = (rr.currentIndex + 1) % len(rr.servers)
	return server, nil
}
