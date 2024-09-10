package main

import (
	"fmt"
	"github.com/zeidlitz/go-balancer/internal/env"
	"github.com/zeidlitz/go-balancer/internal/server"
)

func main() {
	serverList := env.GetString("SERVERS", "server1, server2")
	host := env.GetString("HOST", "localhost")
	port := env.GetString("PORT", "8080")
	host = fmt.Sprint(host + ":" + port)
	server.Start(host)
}
