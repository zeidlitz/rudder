package main

import (
	"fmt"
	"log/slog"

	"github.com/zeidlitz/rudder/internal/env"
	"github.com/zeidlitz/rudder/internal/server"
)

func main() {
	serverlist := env.GetString("SERVERS", "http://34.88.120.114")
	hostname := env.GetString("HOSTNAME", "localhost")
	algorithm := env.GetString("ALGORITHM", "roundrobin")
	port := env.GetString("PORT", "8080")

	server := server.Server{}
	host := fmt.Sprint(hostname + ":" + port)
	server.Configure(algorithm, serverlist)
	slog.Info("Rudder is running!")
	slog.Info("configuration", "serverlist", serverlist, "hostname", hostname, "port", port, "algorithm", algorithm)
	server.Start(host)
}
