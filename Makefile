export CGO_ENABLED = 0

build:
	go build ./cmd/miscutils -ldflags="-w -s"

lint:
	golangci-lint run -c ./golangci.yml ./...
