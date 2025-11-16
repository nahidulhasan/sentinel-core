.PHONY: build run fmt test

build:
	go build ./...

run:
	go run ./cmd/api

fmt:
	gofmt -w .

test:
	go test ./... -v
