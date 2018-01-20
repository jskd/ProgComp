
#include "worksheets.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>


#define MAX_BUF_LINE	(1 << 11)
#define VAL_PER_LINE	(1 << 10)
#define MAX_LINES_COUNT	(1 << 10)

// We're creating arrays so we need fixed sizes
// Unless we introducing another data type !

int parse_data(const char *path, struct worksheet *output) {
	FILE *p_file = NULL;
	char *psz_token = NULL, *p_saveptr = NULL;
	char psz_buf[MAX_BUF_LINE] = {'\0'};
	unsigned u_line_count = 0, u_elem_count = 0;

	if ((p_file = fopen(path, "r")) == NULL)
		return -1;

	struct line_data *pst_lines = malloc(MAX_LINES_COUNT * sizeof(struct line_data));

	while(fgets(psz_buf, sizeof(psz_buf), p_file)) {
		struct line_data *st_current_line = malloc(sizeof(struct line_data));

		psz_buf[strlen(psz_buf)-1] = '\0'; // we're removing the newline
		psz_token = strtok_r(psz_buf, ";", &p_saveptr); // let's parse

		st_current_line->content = malloc(VAL_PER_LINE * sizeof(struct cell));

		// Analyzing each token
		while (psz_token != NULL) {
			struct cell *pst_current_cell = &(st_current_line->content[u_elem_count]);

			switch (token_type(psz_token)) {
				case FORMULA:
					parse_formula(psz_token, pst_current_cell);
					break;
				case VALUE:
					pst_current_cell->ty = VALUE;
					pst_current_cell->udata.value = atoi(psz_token);
					break;
				case INVALID:
				default:
					break;
			}

			// Adding the element
			st_current_line->content[u_elem_count] = *pst_current_cell;

			// let's get to the next token
			psz_token = strtok_r(NULL, ";", &p_saveptr);

			u_elem_count++;
		}

		// Adding the line
		st_current_line->nb_elements = u_elem_count;
		pst_lines[u_line_count] = *st_current_line;

		u_elem_count = 0;
		u_line_count++;
	}

	// Exporting...
	output->nblines = u_line_count;
	output->lines = pst_lines;

	fclose(p_file);
	return 0;
}

int parse_formula(char* psz_token, struct cell *formula) {

    char *formula_element = NULL, *p_saveptr = NULL;
	unsigned nb_parsed_elements = 0;

	formula_element = strtok_r(&(psz_token[2]), ",", &p_saveptr);;

	// Main assumption:
	// A formula contains 5 components

	while (formula_element != NULL) {
		nb_parsed_elements++;

		// If you have a better solution, do it (I'm tired)
		switch (nb_parsed_elements) {
			case 1:
				formula->udata.formula.r1 = atoi(formula_element); break;
			case 2:
				formula->udata.formula.c1 = atoi(formula_element); break;
			case 3:
				formula->udata.formula.r2 = atoi(formula_element); break;
			case 4:
				formula->udata.formula.c2 = atoi(formula_element); break;
			case 5:
				formula->udata.formula.val = atoi(formula_element); break;
			default: break;
		}

		formula_element = strtok_r(NULL, ",", &p_saveptr);
	}

	if (nb_parsed_elements != 5) {
		formula->ty = INVALID;
		return 1;
	}

	formula->ty = FORMULA;

	return 0;
}

int parse_user(const char* path, struct user_data *user_mods) {
	FILE *p_file = NULL;
	char *psz_token = NULL, *p_saveptr = NULL;
	char psz_buf[MAX_BUF_LINE] = {'\0'};
	unsigned nb_parsed_elements = 0, user_changes_count = 0;

	if ((p_file = fopen(path, "r")) == NULL)
		return -1;

	struct user_component *user_changes =
		malloc(MAX_LINES_COUNT * sizeof(struct user_component));

	struct user_component *current_change = user_changes;

	while(fgets(psz_buf, sizeof(psz_buf), p_file)) {
		user_changes_count++;

		psz_buf[strlen(psz_buf)-1] = '\0'; // we're removing the newline
		psz_token = strtok_r(psz_buf, " ", &p_saveptr); // let's parse

		while (psz_token != NULL) {
			nb_parsed_elements++;

			switch (nb_parsed_elements) {
				case 1: current_change->r = atoi(psz_token); break;
				case 2: current_change->c = atoi(psz_token); break;
				case 3:
					if (token_type(psz_token) == FORMULA)
						parse_formula(psz_token, &(current_change->value));
					else {
						current_change->value.udata.value = atoi(psz_token);
						current_change->value.ty = VALUE;
					}
					break;
				default: break;
			}

			psz_token = strtok_r(NULL, " ", &p_saveptr);
		}

		nb_parsed_elements = 0;
		current_change++;
	}

	user_mods->nb_changes = user_changes_count;
	user_mods->content = user_changes;

	return 0;
}

// we should do more checks than just testing the first char
enum cell_ty token_type(const char* psz_token) {

    if (psz_token != NULL) {
		if (psz_token[0] == '=')
			return FORMULA;
		else
			return VALUE;
	}
	return INVALID;
}

void evaluate_worksheet(struct worksheet *ws) {
	printf("Students! This is our job!\n");
}

void apply_user(struct worksheet *ws, struct user_data *user_mods) {
	printf("Students! This is our job!\n");
}

void produce_view(struct worksheet *ws, const char *path) {
	printf("Students! This is our job!\n");
}

void produce_changes(struct worksheet *ws, const char *path) {
	printf("Students! This is our job!\n");
}

void release_worksheet(struct worksheet *ws) {
	for (unsigned i = 0; i < ws->nblines; i++) {
		free(ws->lines[i].content);
	}
	free(ws->lines);
}

void release_user(struct user_data *user_mods) {
	free(user_mods->content);
}

void print_worksheet(struct worksheet *ws) {

    for (unsigned i = 0; i < ws->nblines; i++) {
        struct line_data st_line_data = ws->lines[i];
        for (unsigned j = 0; j < st_line_data.nb_elements; j++) {
			struct cell node = st_line_data.content[j];
			print_data(&node, "\t");
		}
		printf("\n");
	}
}

void print_user(struct user_data *user_mods) {
	for (unsigned i = 0; i < user_mods->nb_changes; i++) {
		struct user_component component = user_mods->content[i];
		printf("(%d;%d) -> ", component.r, component.c);
		print_data(&(component.value), "\n");
	}
}

void print_data(struct cell* value, const char* separator) {
	switch (value->ty) {
		case FORMULA:
			printf("F%s", separator);
			break;
		case VALUE:
			printf("%d%s", value->udata.value, separator);
			break;
		case INVALID:
			printf("?%s", separator);
			break;
		default:
			break;
	}
}
