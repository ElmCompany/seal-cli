#!/usr/bin/make -f
VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"

GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin
GOSRC=$(GOPATH)/src

default: build

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build:
	go build $(LDFLAGS) .

.PHONY: test
test:
	go test -cover  -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html
.PHONY: run
run:
	go run .
