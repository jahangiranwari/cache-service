export GO111MODULE=on

.PHONY: start-server start-worker build-start

build:
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux go build -o ./bin ./...

start-server:
	./bin/cache-service

start-worker:
	./bin/worker

build-start: build start-server
