.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: build


build:
	    @go build $(GOFLAGS) ./...

install:
	    @go get $(GOFLAGS) ./...

test: install
	    @go test $(GOFLAGS) ./...

clean:
	    @go clean $(GOFLAGS) -i ./...
