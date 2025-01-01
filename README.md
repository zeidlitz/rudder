# Rudder

A simple , lightweight HTTP load balancer written in Go

# configuration

Configration is done through environment variables, se below list for supported ones

| Name   | Type   | Description   | Default   |
|---------------- | --------------- | --------------- | --------------- |
| SERVERS   | String  | A list of the servers to load balance between. Seperated with comma (,) | server1, server2 |
| ALGORITHM | String  | The loadbalancing algorithm to use    | roundrobin |
| HOSTNAME  | String  | The host to run the rudder service on | localhost  |
| PORT      | String  | The port to run the rudder service on | 8080       |

# example - roundrobin balacning between two servers

Running this base configuration (a modified version of main.go)

```go
package main

import (
	"fmt"
	"log/slog"

	"github.com/zeidlitz/rudder/internal/env"
	"github.com/zeidlitz/rudder/internal/server"
)

func main() {
    serverlist := env.GetString("SERVERS", "http://34.88.120.114", "http://34.88.120.115")
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

```

This can ofcourse be configured trough env variables. One important thing to note here is that as of now the protocol needs to be supplied on the SERVERS flag, so for example we need to specify HTTP for our two servers


```go
    serverlist := env.GetString("SERVERS", "http://34.88.120.114", "http://34.88.120.115")
```

here we have two servers that will respond to incoming HTTP requests with a 200 OK response, we will run rudder in front of them like so:

```bash
[] ~/repos/common/rudder <main> ✗ go run cmd/rudder.go 
2025/01/01 21:45:21 INFO Rudder is running!
2025/01/01 21:45:21 INFO configuration serverlist=http://34.88.120.114 hostname=localho
st port=8080 algorithm=roundrobin
```

we get confirmation that rudder is running towards the servers we configured. (Note: at this stage we don't do any kind of confirmation that we have connectivity or anything else towards the servers, we just place them in memory and run)

Now if we send a HTTP request towards rudder (in this case on localhost:8080) we it will send it to the first server in the list and relay the response back to us:

```bash
[] ~/repos/common/rudder <main> ✗ curl localhost:8080
{"server":"running","status":"200 OK"}
```

Running the request again will serve it to the next server in the list.
