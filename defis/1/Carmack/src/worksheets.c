#include "worksheets.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_BUF_LINE	(1 << 11)
#define VAL_PER_LINE	(1 << 10)
#define MAX_LINES_COUNT	(1 << 10)

void write_data(FILE* stream, struct cell *value, const char *separator);

// We're creating arrays so we need fixed sizes
// Unless we introducing another data type !
/*
    from @luxon
    1 - !!! Check return values from malloc !!!
    2 - I suggest to use bool instead of int as type of return value, because
        you want to check if the operation has deen done successfully
    3 - Maybe you should put the big while-block in a separate (private) function
*/
int parse_data(const char *path, struct worksheet *output) {
	FILE *p_file = NULL;
	char *psz_token = NULL, *p_saveptr = NULL;
	char psz_buf[MAX_BUF_LINE] = {'\0'};
	unsigned u_line_count = 0, u_elem_count = 0;

	if ((p_file = fopen(path, "r")) == NULL)
		return -1;

	struct line_data *pst_lines = malloc(MAX_LINES_COUNT * sizeof(struct line_data));
	if (pst_lines == NULL)
		return -1;

	/* Suggestion: this block shoud be in a separate function */

	while(fgets(psz_buf, sizeof(psz_buf), p_file)) {
		struct line_data *st_current_line = malloc(sizeof(struct line_data));
		if (st_current_line == NULL)
			return -1;

		psz_buf[strlen(psz_buf)-1] = '\0'; // we're removing the newline
		psz_token = strtok_r(psz_buf, ";", &p_saveptr); // let's parse

		st_current_line->pst_content =
			malloc(VAL_PER_LINE * sizeof(struct cell));
		if (st_current_line->pst_content == NULL)
			return -1;

		// Analyzing each token
		while (psz_token != NULL) {
			struct cell *pst_current_cell =
				&(st_current_line->pst_content[u_elem_count]);

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
			st_current_line->pst_content[u_elem_count] = *pst_current_cell;

			// let's get to the next token
			psz_token = strtok_r(NULL, ";", &p_saveptr);

			u_elem_count++;
		}

		// Adding the line
		st_current_line->nb_elements = u_elem_count;
		pst_lines[u_line_count] = *st_current_line;

		// Resetting/Updating counters
		u_elem_count = 0;
		u_line_count++;
	}

	/* END Suggestion */

	// Exporting...
	output->nblines = u_line_count;
	output->pst_line_data = pst_lines;

	fclose(p_file);
	return 0;
}

int parse_formula(char* psz_token, struct cell *formula) {
    char *psz_formula = NULL, *p_saveptr = NULL;
	unsigned nb_parsed_elements = 0;

	// psz_token+3 because we're skipping `=#(`
	psz_formula = strtok_r((psz_token+3), ",", &p_saveptr);;

	// Main assumption:
	// A formula contains 5 components

	while (psz_formula != NULL) {
		nb_parsed_elements++;

		// If you have a better solution, do it (I'm tired)
		switch (nb_parsed_elements) {
			case 1:
				formula->udata.st_formula.r1 = atoi(psz_formula); break;
			case 2:
				formula->udata.st_formula.c1 = atoi(psz_formula); break;
			case 3:
				formula->udata.st_formula.r2 = atoi(psz_formula); break;
			case 4:
				formula->udata.st_formula.c2 = atoi(psz_formula); break;
			case 5:
				formula->udata.st_formula.val = atoi(psz_formula); break;
			default: break;
		}

		psz_formula = strtok_r(NULL, ",", &p_saveptr);
	}

	// Exporting
	if (nb_parsed_elements != 5) {
		formula->ty = INVALID;
		return 1;
	}

	formula->ty = FORMULA;
    formula->udata.st_formula.status = 0;
    formula->udata.st_formula.res = -1; 
	//IMPORTANT! NOT CALCULATED FORMULAS ARE NOW DEFINED BY -1 IN RES!!!

	return 0;
}

/*
    3 - The "big" while-block in a separate (private) function?
*/
int parse_user(const char* path, struct user_data *user_mods) {
	FILE *p_file = NULL;
	char *psz_token = NULL, *p_saveptr = NULL;
	char psz_buf[MAX_BUF_LINE] = {'\0'};
	unsigned nb_parsed_elements = 0, nb_user_changes = 0;

	if ((p_file = fopen(path, "r")) == NULL)
		return -1;

	struct user_component *pstar_user_changes =
		malloc(MAX_LINES_COUNT * sizeof(struct user_component));
	if (pstar_user_changes == NULL)
		return -1;

	struct user_component *pstar_usr_cur_change = pstar_user_changes;

	while(fgets(psz_buf, sizeof(psz_buf), p_file)) {
		nb_user_changes++;

		psz_buf[strlen(psz_buf)-1] = '\0'; // we're removing the newline
		psz_token = strtok_r(psz_buf, " ", &p_saveptr); // let's parse

		while (psz_token != NULL) {
			nb_parsed_elements++;

			switch (nb_parsed_elements) {
				case 1: pstar_usr_cur_change->r = atoi(psz_token); break;
				case 2: pstar_usr_cur_change->c = atoi(psz_token); break;
				case 3:
					if (token_type(psz_token) == FORMULA)
						parse_formula(psz_token, &(pstar_usr_cur_change->st_value));
					else {
						pstar_usr_cur_change->st_value.udata.value = atoi(psz_token);
						pstar_usr_cur_change->st_value.ty = VALUE;
					}
					break;
				default: break;
			}

			// Let's get to the next token
			psz_token = strtok_r(NULL, " ", &p_saveptr);
		}

		// Resetting/Updating counters
		nb_parsed_elements = 0;
		pstar_usr_cur_change++;
	}

	// Exporting
	user_mods->nb_changes = nb_user_changes;
	user_mods->pst_content = pstar_user_changes;

	fclose(p_file);
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

//struct user_data log[MAX_LINES_COUNT];

int evaluate_formula(struct formula *f, struct worksheet *ws,
                     int x0, int y0) {
    int x, y, val_num = 0;
    struct cell *c;

    if(f->status && f->res == -1) //visited and not calculated formula -> error
        goto return_with_error;
	else{
		if (f->status) //supposed to be well calculated formula
			return 0;
	}
    f->status = 1;

    for(x = f->r1; x <= f->r2; x++) {
        for(y = f->c1; y <= f->c2; y++) {
            c = (ws->pst_line_data + x)->pst_content + y;
            switch(c->ty) {
            case INVALID:
                goto return_with_error;
            case VALUE:
                val_num = (c->udata.value == f->val) ?
                    (val_num + 1) : val_num;
                break;
            case FORMULA:
				if (c->udata.st_formula.res == -1){ 
				//if not INVALID and res == -1 then not yet calculated
                // Look, Ma, recursion in C!
                evaluate_formula(&(c->udata.st_formula), ws, x, y);
                y--;
				}
				else {
					val_num = (c->udata.st_formula.res == f->val) ?
                    (val_num + 1) : val_num;
				}
                break;
            }
        }
    }

    //((ws->pst_line_data + x0)->pst_content + y0)->ty = VALUE; 
	//formula stays formula
    ((ws->pst_line_data + x0)->pst_content + y0)->udata.st_formula.res = val_num;
    return 0;

 return_with_error:
    ((ws->pst_line_data + x0)->pst_content + y0)->ty = INVALID;
    return -1;
}

void evaluate_worksheet(struct worksheet *ws) {
    int x, y;
    struct cell *c;

    for(x = 0; x < ws->nblines; x++) {
        for(y = 0; y < (ws->pst_line_data + x)->nb_elements; y++) {
            c = (ws->pst_line_data + x)->pst_content + y;
            if(c->ty == FORMULA)
                evaluate_formula(&(c->udata.st_formula), ws, x, y);
        }
    }
}

int produce_view(struct worksheet *ws, const char *path) {
	FILE *p_file = fopen(path, "w");
	unsigned i, j;

	if (p_file == NULL)
		return -1;

	for (i = 0; i < ws->nblines; i++) {
        struct line_data st_line_data = ws->pst_line_data[i];
		unsigned nb_elements = st_line_data.nb_elements;
        for (j = 0; j < nb_elements; j++) {
			struct cell node = st_line_data.pst_content[j];
			write_data(p_file, &node, (j == nb_elements-1) ? "" : ";");
		}
		fprintf(p_file, "\n");
	}

	return 0;
}

void write_change_prologue(FILE* stream, struct user_component mod) {
	switch (mod.st_value.ty) {
		case VALUE:
			fprintf(stream, "after \"%d %d %d\":\n",
					mod.r,
					mod.c,
					mod.st_value.udata.value);
			break;
		case FORMULA:
			if (mod.st_value.udata.st_formula.val != '?') {
				fprintf(stream, "after \"%d %d =#(%d,%d,%d,%d,%d)\":\n",
						mod.r,
						mod.c,
						mod.st_value.udata.st_formula.r1,
						mod.st_value.udata.st_formula.r2,
						mod.st_value.udata.st_formula.c1,
						mod.st_value.udata.st_formula.c2,
						mod.st_value.udata.st_formula.val);
			}
			break;
		case INVALID: // Do nothing
		default:
			break;
	}
}

void write_change(FILE* stream, int line, int col, struct cell value, char verbose) {
	int val;
	switch (value.ty) {
		case INVALID:
			val = -1;
			break;
		case VALUE:
			val = value.udata.value;
			break;
		case FORMULA:
			val = value.udata.st_formula.res;
			break;
		default:
			val = -1;
			break;
	}
	if (verbose)
		printf("(%d,%d) %d -> %d\n", line, col, value.old_value, val);
	fprintf(stream, "%d %d %d\n", line, col, val);
}

void reset_table(struct worksheet *ws){
	//added old_value to all cells, super memory-consumung, redo later with log
	//after this function all old_values are set correctly
	unsigned i, j;
	for (i = 0; i < ws->nblines; i++) {
    struct line_data st_line_data = ws->pst_line_data[i];
	unsigned nb_elements = st_line_data.nb_elements;
        for (j = 0; j < nb_elements; j++) {
			struct cell *node = st_line_data.pst_content + j;
			switch(node->ty){
				case FORMULA:
					//if formula survived evaluate_formula, it has good res
					node->old_value = node->udata.st_formula.res;
					node->udata.st_formula.res = -1;
					node->udata.st_formula.status = 0;
					break;
				case VALUE:
					node->old_value = node->udata.value;
					break;
				case INVALID:
/*
supposing node->udata.st_formula.res was left -1 after evaluate_formula,
if not the case, OBLIGATORILY uncomment next line
node->udata.st_formula.res = -1
*/					
					node->old_value = -1;
					node->ty = FORMULA;
					node->udata.st_formula.status = 0;
					break;
				default:
					break;
			}
		}
	}
}

int produce_changes(struct worksheet *ws,  struct user_data *user_mods, const char *path) {
	int changes_counter = 0;

	reset_table(ws);
	apply_user(ws, user_mods);

	FILE *p_file = fopen(path, "w");
	
	if (p_file == NULL)
		return -1;
	
	unsigned i, j;
	for (i = 0; i < ws->nblines; i++) {
        struct line_data st_line_data = ws->pst_line_data[i];
		unsigned nb_elements = st_line_data.nb_elements;
        for (j = 0; j < nb_elements; j++) {
			struct cell node = st_line_data.pst_content[j];
			if ((node.ty == INVALID && node.old_value != -1) ||
			(node.ty == VALUE && node.old_value != node.udata.value) ||
			(node.ty == FORMULA && node.old_value != node.udata.st_formula.res))
			{
				write_change_prologue(p_file,
						user_mods->pst_content[changes_counter]);
				write_change(p_file, i, j, node, 0); //change verbose to 0 here
				changes_counter++;
			}
		}
		//fprintf(p_file, "\n");
	}

	return 0;
}

void release_worksheet(struct worksheet *ws) {
	for (unsigned i = 0; i < ws->nblines; i++) {
		free(ws->pst_line_data[i].pst_content);
	}
	free(ws->pst_line_data);
}

void release_user(struct user_data *user_mods) {
	free(user_mods->pst_content);
}

void print_worksheet(struct worksheet *ws) {
    for (unsigned i = 0; i < ws->nblines; i++) {
        struct line_data st_line_data = ws->pst_line_data [i];
        for (unsigned j = 0; j < st_line_data.nb_elements; j++) {
			struct cell node = st_line_data.pst_content[j];
			write_data(stdout, &node, "\t");
		}
		printf("\n");
	}
	printf("\n");
}

//don't use it for a while
void print_user(struct user_data *user_mods) {
	for (unsigned i = 0; i < user_mods->nb_changes; i++) {
		struct user_component component = user_mods->pst_content[i];
		printf("(%d;%d) -> ", component.r, component.c);
		write_data(stdout, &(component.st_value), "\n");
	}
}

void write_data(FILE* stream, struct cell* value, const char* separator) {
	switch (value->ty) {
		case FORMULA:
			if (value->udata.st_formula.res == -1) //not yet calculated formula
				fprintf(stream, "F%s", separator);
			else 
				fprintf(stream, "%d%s", value->udata.st_formula.res, separator);
			break;
		case VALUE:
			fprintf(stream, "%d%s", value->udata.value, separator);
			break;
		case INVALID:
			fprintf(stream, "P%s", separator);
			break;
		default:
			break;
	}
}


void apply_user(struct worksheet *ws, struct user_data *user_mods) {
    int i;
    struct user_component c;

    for(i = 0; i < user_mods->nb_changes; i++) {
        c = user_mods->pst_content[i];
        if(c.r < 0 || c.r > ws->nblines ||
           c.c < 0 || c.c > ws->pst_line_data[c.r].nb_elements) {
            fprintf(stderr, "Out of bounds!\n");
            continue;
        }
		struct cell *target = ws->pst_line_data[c.r].pst_content + c.c;
        target->ty = c.st_value.ty;
		target->udata = c.st_value.udata; //should work fine 
    }
    evaluate_worksheet(ws);
}
