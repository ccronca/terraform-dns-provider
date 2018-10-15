PKGS := $(shell go list ./... | grep -v /vendor)

# Go parameters
GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BIN_DIR := $(GOPATH)/bin

all: deps test

test:
	$(GOTEST) -v -timeout 120s -short $(PKGS)

clean:
	$(GOCLEAN)

deps:
	$(GOGET) -u github.com/golang/dep/cmd/dep
	dep ensure -v

.PHONY: all test deps
