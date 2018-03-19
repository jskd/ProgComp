CC     = g++
CFLAGS = -DCHICKEN_MODE -Wall -O2 -std=c++11

OBJS   = rtree.o formulanode.o area.o parser.o formula.o

ifeq ($(OS),Windows_NT)
	LIB = C:\Windows\System32\OpenCL.dll
	WS     = ws.exe
else
	LIB = -lOpenCL
	WS     = ws
endif

$(WS): $(OBJS)
	$(CC) $(CFLAGS) main.cpp -o $(WS) $(STATIC) $(OBJS) $(LIB)

%.o: %.c %.h
	gcc $(CFLAGS) -c $< -o $@

%.o: %.cpp %.hpp
	$(CC) $(CFLAGS) -c $< -o $@

clean:; rm -rf $(OBJS)

mrproper: clean
	rm -rf $(WS)

static: STATIC := -static 
static: $(WS) 
	