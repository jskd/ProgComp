CC     = g++
CFLAGS = -DCHICKEN_MODE -Wall -O2 -std=c++11

WS     = ws
OBJS   = rtree.o formulanode.o area.o parser.o formula.o

$(WS): $(OBJS)
	$(CC) $(CFLAGS) main.cpp -o $(WS) $(OBJS) -lOpenCL

%.o: %.c %.h
	gcc $(CFLAGS) -c $< -o $@

%.o: %.cpp %.hpp
	$(CC) $(CFLAGS) -c $< -o $@

clean:; rm -rf $(OBJS)

mrproper: clean
	rm -rf $(WS)
