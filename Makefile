.PHONY: all test clean build install

GOFLAGS ?= $(GOFLAGS:)

all: install deps


install:
	    @go install -v $(GOFLAGS) ./...

deps:
	    @go get -v $(GOFLAGS) ./...

test: install
	    @go test $(GOFLAGS) ./...

clean:
	    @go clean $(GOFLAGS) -i ./...
