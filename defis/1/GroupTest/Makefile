CPP=g++ --std=c++11 -Wall
EXEC = ws

EXEC_FILE= $(EXEC)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	LIBFLAGS =
endif
ifeq ($(UNAME_S),Darwin)
	LIBFLAGS =
endif

all: $(EXEC_FILE)

$(EXEC_FILE): main.cpp
	$(CPP) -o $@ $^ $(LIBFLAGS)

run:
	./$(EXEC_FILE)

clean:
	rm -rf *.o

mrproper: clean
	rm -rf ws
