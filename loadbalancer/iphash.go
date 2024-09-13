package loadbalancer

type IPHash struct {
	servers []string
}

func (iph *IPHash) Configure(serverlist []string) error {
	iph.servers = serverlist
	return nil
}

func (iph *IPHash) GetServer() (string, error) {
	// TOD: Implement this
	return iph.servers[0], nil
}
