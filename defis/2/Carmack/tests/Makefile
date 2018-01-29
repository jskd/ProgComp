CC		= g++
CFLAGS	= -g -std=c++11 -Wall -pedantic
OBJECTS	= main.o
TARGET	= wst

all: $(TARGET)

$(TARGET): $(OBJECTS)
	$(CC) $^ -o $(TARGET) $(CFLAGS)

main.o : main.cpp
	$(CC) -c $< -o $@ $(CFLAGS)

clean:
	rm -rf *.o

mrproper: clean
	rm -rf $(TARGET)

