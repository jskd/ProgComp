SHELL = /bin/sh

DOCKER = docker

IMG = nw-group/ws

.PHONY: all build-docker-img mrproper

all: build-docker-img

build-docker-img:
	$(DOCKER) build -t $(IMG) .

mrproper:
	$(DOCKER) rmi -f $(IMG) 2> /dev/null
