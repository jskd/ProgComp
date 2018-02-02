CPP=g++ --std=c++11 -Wall
EXEC = ws

SRC_DIR = src
BIN_DIR = ./
EXEC_FILE= $(EXEC)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	LIBFLAGS =
endif
ifeq ($(UNAME_S),Darwin)
	LIBFLAGS =
endif

SRC_FILES= $(wildcard $(SRC_DIR)/**.cpp) $(wildcard $(SRC_DIR)/**/**.cpp) $(wildcard $(SRC_DIR)/**/**/**.cpp) $(wildcard $(SRC_DIR)/**/**/**/**.cpp)
OBJ_FILES= $(patsubst %.cpp, %.o,  $(subst $(SRC_DIR), $(BIN_DIR),$(SRC_FILES)))

all: $(EXEC_FILE)

$(EXEC_FILE): $(OBJ_FILES)
	$(CPP) -o $@ $^ $(LIBFLAGS)

$(OBJ_FILES): %.o : $(SRC_DIR)/%.cpp
	@mkdir -p $(dir $@)
	$(CPP) -c -o $@ $^

run:
	./$(EXEC_FILE)

clean:
	rm -rf *.o

mrproper: clean
	rm -rf ws
