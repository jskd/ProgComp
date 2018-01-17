CC=gcc
CFLAGS=-W -Wall -ansi -pedantic
TARGET=ws

$(TARGET): main.c
	$(CC) main.c -o $(TARGET) $(CFLAGS)

clean:
		rm -rf *.o

mrproper: clean
		rm -rf $(TARGET)
