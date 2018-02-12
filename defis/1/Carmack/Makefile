CC     = gcc
CFLAGS = -g -std=gnu99 -Wall -pedantic
OBJECTS = src/main.o src/worksheets.o
TARGET = wso
DOCKER_IMAGE = carmack-image

all: docker_alpine

docker_alpine:
	@docker build -f Dockerfile_alpine -t $(DOCKER_IMAGE) .

docker_gcc:
	@docker build -f Dockerfile_gcc -t $(DOCKER_IMAGE) .

fresh: mrproper $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $^ -o $(TARGET) $(CFLAGS)

src/worksheets.o : src/worksheets.c src/worksheets.h
	$(CC) -c $< -o $@ $(CFLAGS)

src/main.o : src/main.c src/worksheets.h
	$(CC) -c $< -o $@ $(CFLAGS)

clean:
	rm -rf src/*.o

mrproper: clean
	rm -rf $(TARGET) *.csv *.txt

.PHONY: all docker_alpine docker_gcc clean mrmroper fresh
