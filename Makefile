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

.PHONY: run
run:
	go run .

# docs swager gen : https://github.com/mayur-tolexo/demo/blob/master/Makefile