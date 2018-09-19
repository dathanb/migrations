.PHONY: all build run-dev

GO_FILES = $(shell find main.go api cli config db handler vendor -type f | sort)

all: build

run-dev: build
	bin/local_env.sh ./migration-demo start

build: migration-demo

migration-demo: $(GO_FILES)
	go build

test: build
	go test -v ./...

docker:
	docker build -t migration-demo .
