

UID:=$(shell id -u)
GID:=$(shell id -g)
PWD:=$(shell pwd)
DIRNAME:=$(shell basename "$(PWD)")


GO_CONTAINER_IMG_NAME=$(DIRNAME)_go
GO_DEFAULT_PORTS=-p 8087:8087
GO_OPTS=-v $(PWD):/go/src/$(DIRNAME) -w /go/src/$(DIRNAME)


.PHONY: cli
cli: ##@setup set up a docker container with mounted source where you can execute all go commands
	docker run -it --rm $(GO_OPTS) $(GO_DEFAULT_PORTS) golang:1.12.6 bash

