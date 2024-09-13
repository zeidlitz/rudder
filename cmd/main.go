package main

import (
	"fmt"
	"github.com/zeidlitz/go-balancer/internal/env"
	"github.com/zeidlitz/go-balancer/internal/server"
)

func main() {
	server := server.Server{}
	serverlist := env.GetString("SERVERS", "server1, server2")
	hostname := env.GetString("HOSTNAME", "localhost")
	algorithm := env.GetString("ALGORITHM", "roundrobin")
	port := env.GetString("PORT", "8080")
	host := fmt.Sprint(hostname + ":" + port)
	server.Configure(algorithm, serverlist)
	server.Start(host)
}
