package main

import (
	"fmt"
	"github.com/zeidlitz/rudder/internal/env"
	"github.com/zeidlitz/rudder/internal/server"
)

func main() {
	serverlist := env.GetString("SERVERS", "server1, server2")
	hostname := env.GetString("HOSTNAME", "localhost")
	algorithm := env.GetString("ALGORITHM", "roundrobin")
	port := env.GetString("PORT", "8080")

	server := server.Server{}
	host := fmt.Sprint(hostname + ":" + port)
	server.Configure(algorithm, serverlist)
	server.Start(host)
}
