# Docker security demo
CSS demo to show how Docker can be used securely. This demo has a GO binary with command injection.

## GO compile vuln app
0. `go build go_http_server.go`

## Run command injection
0. `curl http://http://127.0.0.1:5000/cmd?key=ls`
0. `curl http://http://127.0.0.1:5000/cmd?key=whoami`


## Spin up vulnerable container
0. `docker-compose -f docker-compose-vuln.yml up -d`

## Spin up NON-vulnerable container
0. `docker-compose -f docker-compose-novuln.yml up -d`

## Resources/Sources
* [Install Go for NATS](https://nats.io/documentation/tutorials/go-install/)
* [Getting Started](https://golang.org/doc/install#testing)
* [Create the smallest and secured golang docker image based on scratch](https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324)
* [Get a URL Parameter from a Request](https://golangcode.com/get-a-url-parameter-from-a-request/)
* [Building a (Web) HTTP Server with Go](https://itnext.io/building-a-web-http-server-with-go-6554029b4079)
