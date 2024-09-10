package main

import (
	"fmt"
	"github.com/zeidlitz/go-balancer/internal/env"
	"github.com/zeidlitz/go-balancer/internal/server"
)

func main() {
	serverList := env.GetString("SERVERS", "server1, server2")
	hostname := env.GetString("HOSTNAME", "localhost")
	port := env.GetString("PORT", "8080")
	host := fmt.Sprint(hostname + ":" + port)
	server.Start(host)
}
