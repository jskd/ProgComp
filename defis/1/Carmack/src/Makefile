CC     = gcc
CFLAGS = -g -std=gnu99 -Wall -pedantic
OBJECTS = main.o worksheets.o
TARGET = ws

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $^ -o $(TARGET) $(CFLAGS)

worksheets.o : worksheets.c worksheets.h
	$(CC) -c $< -o $@ $(CFLAGS)

main.o : main.c worksheets.h
	$(CC) -c $< -o $@ $(CFLAGS)

clean:
	rm -rf *.o

mrproper: clean
		rm -rf $(TARGET)
