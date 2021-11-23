#!/usr/bin/make -f
VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"

GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin
GOSRC=$(GOPATH)/src

default: build

.PHONY: build
build:
	go build $(LDFLAGS) .

.PHONY: test
test:
	go test -cover  -v ./...

.PHONY: run
run:
	go run .
