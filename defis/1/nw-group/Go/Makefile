GOPATH:=$(shell pwd)

all: export_gopath build

export_gopath:
	export GOPATH=$(GOPATH)

build:
	go build -o main.exe src/*.go

test:
	go test ./... -v

clean:
	go clean
