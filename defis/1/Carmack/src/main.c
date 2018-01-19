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
	user current_user;

	parse_data(argv[1], &current_worksheet);
	parse_user(argv[2], &current_user);
	print_worksheet(&current_worksheet);
	print_user(&current_user);
	release_worksheet(&current_worksheet);
	release_user(&current_user);

	return EXIT_SUCCESS;
}

