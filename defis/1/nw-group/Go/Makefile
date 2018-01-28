GOPATH:=$(shell pwd)

all: export_gopath build

export_gopath:
	export GOPATH=$(GOPATH)

build:
	go build -o main.exe ./*.go

test:
	go test ./... -v
