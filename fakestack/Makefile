.PHONY: all build run-dev

GO_FILES=$(shell find main.go api cli config db handler vendor -type f | sort)
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
	docker build -t fakestack .
	docker tag fakestack:latest fakestack:$(VERSION)
