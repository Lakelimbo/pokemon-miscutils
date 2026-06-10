export CGO_ENABLED=0

build:
	go build -o ./build/ -ldflags="-w -s" ./cmd/miscutils

lint:
	golangci-lint run -c ./golangci.yml ./...
