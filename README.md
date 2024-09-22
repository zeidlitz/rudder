# Rudder

A simple , lightweight HTTP load balancer written in Go

# configuration

Configration is done through environment variables, se below list for supported ones

| Name   | Type   | Description   | Default   |
|---------------- | --------------- | --------------- | --------------- |
| SERVERS   | String  | A list of the servers to load balance between. Seperated with comma (,) | server1, server2 |
| ALGORITHM | String  | The loadbalancing algorithm to use     | roundrobin |
| HOSTNAME  | String  | The host to run the rudder service on | localhost  |
| PORT      | String  | The port to run the rudder service on | 8080       |
