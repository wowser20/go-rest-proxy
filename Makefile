include .env

# set default shell
SHELL = bash -e -o pipefail

# variables
VERSION                  ?= $(shell cat ./VERSION)

now=$(shell date +"%Y%m%d%H%M%S")

default: run

.PHONY:	install
install:
	go mod tidy
	go mod vendor

.PHONY:	lint
lint:
	golangci-lint run 

.PHONY:	build
build:
	mkdir -p bin
	go build -race -o bin/${API_NAME} \
	    cmd/main.go

.PHONY:	test
test:
	go test -race -v -p 1 ./...
	
.PHONY:	run
run:	build
	./bin/${API_NAME}

.PHONY: up
up:
	docker compose down
	docker compose up -d --build