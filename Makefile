.PHONY: all

GO_FILES = $(shell find api cli db error config vendor -type f | sort)

all: migration-demo

migration-demo: $(GO_FILES)
	go build

test: build
	go test -v ./...
