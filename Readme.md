# Demo Load Balancer in GO

This is a demo load balancer written in GO. It is a simple load balancer that uses a round robin algorithm to distribute the load across the servers.

## How to run

1. Clone the repo
2. Run `go run main.go` to start the load balancer

## How to test

I run 3 servers with `npx http-server` and then run the load balancer. I then run `curl localhost:8080` and it will return the response from the servers.

![screenshot-image](https://i.imgur.com/gWuLoY0.png)