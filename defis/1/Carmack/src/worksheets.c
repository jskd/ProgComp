#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "worksheets.h"

#define MAX_BUF_LINE	(1 << 11)
#define VAL_PER_LINE	(1 << 10)
#define MAX_LINES_COUNT	(1 << 10)

// We're creating arrays so we need fixed sizes
// Unless we introducing another data type !

int parse_data(const char *path, worksheet* output) {
	FILE *file;
	char *token, *saveptr;
	char buf[MAX_BUF_LINE];
	unsigned lines_count = 0, elems_count = 0;
	
	file = fopen(path, "r");

	if ((file = fopen(path, "r")) == NULL)
		return 1;

	line *lines = malloc(MAX_LINES_COUNT * sizeof(line));

	while(fgets(buf, sizeof(buf), file)) {
		line *current_line = malloc(sizeof(line));

		buf[strlen(buf)-1] = '\0'; // we're removing the newline
		token = strtok_r(buf, ";", &saveptr); // let's parse

		current_line->content = malloc(VAL_PER_LINE * sizeof(data));
		
		while (token != NULL) { 
			data *current_data = &(current_line->content[elems_count]);

			switch (token_type(token)) {
				case FORMULA:
					parse_formula(token, current_data);
					break;
				case VALUE:
					current_data->ty = VALUE;
					current_data->value = atoi(token);
					break;
				case INVALID:
				default:
					break;
			}

			// Adding the element
			current_line->content[elems_count] = *current_data;

			// let's get to the next token
			token = strtok_r(NULL, ";", &saveptr);

			elems_count++;
		}
		
		// Adding the line
		current_line->elements_count = elems_count;
		lines[lines_count] = *current_line;

		elems_count = 0;
		lines_count++;
	}

	output->lines_count = lines_count;
	output->lines = lines;

	fclose(file);

	return 0;
}

int parse_formula(char* token, data* formula) {
	char* formula_element, *saveptr;
	unsigned elements_parsed = 0;

	formula_element = strtok_r(&(token[2]), ",", &saveptr);;

	while (formula_element != NULL) {
		elements_parsed++;

		// If you have a better solution, do it (I'm tired)
		switch (elements_parsed) {
			case 1:
				formula->formula.r1 = atoi(formula_element);
				break;
			case 2:
				formula->formula.c1 = atoi(formula_element);
				break;
			case 3:
				formula->formula.r2 = atoi(formula_element);
				break;
			case 4:
				formula->formula.c2 = atoi(formula_element);
				break;
			case 5:
				formula->formula.val = atoi(formula_element);
				break;
			default:
				break;
		}

		formula_element = strtok_r(NULL, ",", &saveptr);
	}

	if (elements_parsed != 5) {
		formula->ty = INVALID;
		return 1;
	}

	formula->ty = FORMULA;

	return 0;	
}

// we should do more checks than just testing the first char
data_ty token_type(const char* token) {
	if (token) {
		if (token[0] == '=')
			return FORMULA;
		else
			return VALUE;
	}
	return INVALID; 
}

void release_worksheet(worksheet* ws) {
	for (unsigned i = 0; i < ws->lines_count; i++) {
		free(ws->lines[i].content);
	}
	free(ws->lines);
}

void print_worksheet(worksheet* ws) {
	for (unsigned i = 0; i < ws->lines_count; i++) {
		line line = ws->lines[i];
		for (unsigned j = 0; j < line.elements_count; j++) {
			data node = line.content[j];
			switch (node.ty) {
				case FORMULA:
					printf("F\t");
					break;
				case VALUE:
					printf("%d\t", node.value);
					break;
				case INVALID:
					printf("?\t");
					break;
				default:
					break;
			}
		}
		printf("\n");
	}
}

