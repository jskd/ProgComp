#include <stdlib.h>
#include <stdio.h>
#include <errno.h>

typedef enum { FORMULA, VALUE } data_ty;

typedef struct {
    int r1, r2, c1, c2, val;
} formula;

typedef struct {
    data_ty ty;
    union {
        int value;
        formula formula;
    };
} data;

int main(int argc, const char* argv[]) {
	if (argc != 5) {
		errno = EINVAL;
		perror("usage: ./ws data.csv user.txt view0.csv changes.txt\n");
	}
	return EXIT_SUCCESS;
}

