.PHONY: all

GO_FILES = $(shell find main.go api cli config db error handler vendor -type f | sort)

all: migration-demo

migration-demo: $(GO_FILES)
	go build

test: build
	go test -v ./...
