## Running locally

* `$ go mod download`
* `$ go run main.go`

## Building container

* `$ docker buildx build -o type=docker --platform=linux/amd64 --tag thedevelopnik/tp-headers-instrumentation-go:1.0 .`
