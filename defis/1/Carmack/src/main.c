#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include "worksheets.h"

int main(int argc, const char* argv[]) {
	if (argc != 5) {
		errno = EINVAL;
		perror("usage: ./ws data.csv user.txt view0.csv changes.txt\n");
		return EXIT_FAILURE;
	}

	worksheet current_worksheet;

	parse_data(argv[1], &current_worksheet);
	print_worksheet(&current_worksheet);
	release_worksheet(&current_worksheet);

	return EXIT_SUCCESS;
}

