CC     = gcc
CFLAGS = -g -std=gnu99 -Wall -pedantic
OBJECTS = src/main.o src/worksheets.o
TARGET = ws


all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $^ -o $(TARGET) $(CFLAGS)

src/worksheets.o : src/worksheets.c src/worksheets.h
	$(CC) -c $< -o $@ $(CFLAGS)

src/main.o : src/main.c src/worksheets.h
	$(CC) -c $< -o $@ $(CFLAGS)

clean:
	rm -rf src/*.o

mrproper: clean
		rm -rf $(TARGET)
