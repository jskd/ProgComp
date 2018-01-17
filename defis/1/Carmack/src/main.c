#include <stdlib.h>
#include <stdio.h>
#include <errno.h>

int main(int argc, const char* argv[]) {
	if (argc != 5) {
		errno = EINVAL;
		perror("usage: ./ws data.csv user.txt view0.csv changes.txt\n");
	}
	return EXIT_SUCCESS;
}

