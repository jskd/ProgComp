CC=sbt
LANG=scala
JAVA=java


all: package

compile: ENV_VERIFICATION
	$(CC) compile

package: ENV_VERIFICATION
	$(CC) assembly

ENV_VERIFICATION:

ifeq (, $(shell which $(JAVA)))
	$(error "Java not found : apt-get install java")
endif

ifeq (, $(shell which $(LANG)))
	$(error "Scala not found : apt-get install scala")
endif

ifeq (, $(shell which $(CC)))
	$(error "SBT not found : apt-get install sbt")
endif

clean:
	rm -rf target
	rm -rf project/project
	rm -rf project/target

mrproper: clean
	rm -rf out/

.PHONY: compile package mrproper clean ENV_VERIFICATION