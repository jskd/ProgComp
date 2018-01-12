anagram: anagram.c linked_list.o multiset.o
	gcc -Wall -o anagram anagram.c linked_list.o multiset.o

%.o: %.c %.h
	gcc -Wall -c $< -o $@

clean:; rm -fr *.o anagram
