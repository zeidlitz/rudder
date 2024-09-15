package loadbalancer

import(
   "crypto/sha256"
)

type IPHash struct {
	servers []string
}

func (iph *IPHash) Configure(serverlist []string) error {
	iph.servers = serverlist
	return nil
}

func (iph *IPHash) GetServer() (string, error) {
  hash := sha256.New()
  hash.Write([]byte(iph.servers[0]))
	return iph.servers[0], nil
}
