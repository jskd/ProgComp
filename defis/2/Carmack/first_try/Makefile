test: test.cpp rtree.cpp cell.cpp area.cpp parser.cpp formula.o
	g++ -o test test.cpp rtree.cpp cell.cpp parser.cpp area.cpp \
	formula.o -lOpenCL

formula.o: formula.c
	gcc -c formula.c

clean:; rm -rf test formula.o
