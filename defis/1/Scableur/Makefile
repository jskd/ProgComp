CC=sbt
LANG=scala
DOCKER=docker

IMG="scableur-img"

.PHONY: compile package mrproper clean build-img


all: build-img

native : package 

compile:
	$(CC) compile

package:
	$(CC) assembly

build-img:
	$(DOCKER) build -t $(IMG) .

clean:
	rm -rf target
	rm -rf project/project
	rm -rf project/target

mrproper: clean
	rm -rf out/

