.PHONY: all build run-dev

GO_FILES=$(shell find . -name '*.go' -type f | sort)
VERSION=$(shell git describe)

all: build

run-dev: build
	bin/local_env.sh ./fakestack start

build: fakestack

fakestack: $(GO_FILES)
	go build

test: build
	go test -v ./...

docker:
	go mod download
	docker build -t fakestack .
	docker tag fakestack:latest fakestack:$(VERSION)
