.PHONY: start_db build

build:
	go build

start_db:
	bin/start-db

test:
	go test -v ./...
