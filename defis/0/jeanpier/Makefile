
CC=g++
FLAGS=-O3 -Wall -Wextra -std=c++11


all: anagram

anagram: anagram.o
	$(CC) $(FLAGS) $< -o $@

anagram.o: anagram.cpp
	$(CC) -c $(FLAGS) $< -o $@

mrproper:
	rm -f anagram *.o
