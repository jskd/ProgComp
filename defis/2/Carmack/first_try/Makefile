test: test.cpp rtree.cpp formulanode.cpp area.cpp parser.cpp formula.o
	g++ -O3 -g -o test test.cpp rtree.cpp formulanode.cpp parser.cpp area.cpp \
	formula.o -lOpenCL

formula.o: formula.c
	gcc -O3 -c -g formula.c

clean:; rm -rf test formula.o
