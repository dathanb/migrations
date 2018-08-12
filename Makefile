.PHONY: all build run-dev

GO_FILES = $(shell find main.go api cli config db error handler vendor -type f | sort)

all: build

run-dev: build
	bin/local-env.sh ./migration-demo

build: migration-demo

migration-demo: $(GO_FILES)
	go build

test: build
	go test -v ./...
