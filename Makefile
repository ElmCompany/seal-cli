#!/usr/bin/make -f

default: build
build:
	go build .

run:
	go run .

# docs swager gen : https://github.com/mayur-tolexo/demo/blob/master/Makefile