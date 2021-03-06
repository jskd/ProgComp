#include "worksheets.h"

#include <stdio.h>
#include <stdlib.h>

#define ARGN 5

int main(int argc, const char* argv[]) {
	if (argc != ARGN) {
		fprintf(stderr, "usage: ./ws data.csv user.txt view0.csv changes.txt\n");
		return EXIT_FAILURE;
	}

	struct worksheet current_worksheet;
	struct user_data current_user;

	if (parse_data(argv[1], &current_worksheet) != 0) {
		fprintf(stderr, "Error while parsing `data.csv` file");
		return EXIT_FAILURE;
	}

	if (parse_user(argv[2], &current_user) != 0) {
		fprintf(stderr, "Error while parsing `user.txt` file");
		return EXIT_FAILURE;
	}


	print_worksheet(&current_worksheet);
	evaluate_worksheet(&current_worksheet);
	print_worksheet(&current_worksheet);

	if (produce_view(&current_worksheet, argv[3])) {
		fprintf(stderr, "Error while writing to `%s` file", argv[3]);
	}
	
	if (produce_changes(&current_worksheet, &current_user, argv[4])) {
		 fprintf(stderr, "Error while writing to  i%s` file", argv[4]);
	}
	
	print_worksheet(&current_worksheet);

	release_worksheet(&current_worksheet);
	release_user(&current_user);

	return EXIT_SUCCESS;
}

