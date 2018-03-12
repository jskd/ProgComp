CC=sbt
LANG=scala
DOCKER=docker
 
IMG="scableur-img"

all: image

compile:
	$(CC) compile

package:
	$(CC) assembly

image:
	$(DOCKER) build -t $(IMG) .

clean:
	rm -rf target
	rm -rf project/project
	rm -rf project/target

mrproper: clean
	rm -rf out/

.PHONY: compile package mrproper clean