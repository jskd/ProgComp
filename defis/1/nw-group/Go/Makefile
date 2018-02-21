SHELL = /bin/sh
RM = rm -f

GOPATH:=$(shell pwd)
EXEC = main.exe

.PHONY: all export_gopath build test clean

all: export_gopath build test

export_gopath:
	export GOPATH=$(GOPATH)

build:
	go build -o $(EXEC) src/main.go

test:
	go test ./... -v

clean:
	$(RM) $(EXEC)
