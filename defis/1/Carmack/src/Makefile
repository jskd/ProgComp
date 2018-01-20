CC     = gcc
CFLAGS = -std=c11 -Wall -pedantic
TARGET = ws

all: $(TARGET)

$(TARGET): *.c
	$(CC) *.c -o $(TARGET) $(CFLAGS)

clean:
	rm -rf *.o

mrproper: clean
		rm -rf $(TARGET)
