/*
    worksheet.h

    Every structures used in the program to execute
    the different operations
*/

#ifndef WORKSHEETS_H
#define WORKSHEETS_H

// Data types

// Some data types may be unnecessary for the moment
// but keep in mind that it is a naive implementation.
// Refactor if you want to but please do not loose semantic value !
// Besides, if we implement a big data worksheet engine in the future,
// we may need them for storing multiple metadata.

// Copy! (^_^)

enum cell_ty { INVALID, FORMULA, VALUE };

struct formula {
	int r1, r2, c1, c2, val, status;
};

/*
    A cell in the worksheet
*/
struct cell {
	enum cell_ty ty;
	union {
		int value;
		struct formula st_formula;
	} udata;
};

struct line_data {
	unsigned nb_elements;
	struct cell *pst_content;
};

struct worksheet {
	unsigned nblines;
	struct line_data *pst_line_data;
};

struct user_component {
	int r, c;
	struct cell st_value;
};

struct user_data {
	unsigned nb_changes;
	struct user_component *pst_content;
};

// Functions

/// Parsing
int parse_data(const char *path, struct worksheet *ws);
int parse_formula(char *token, struct cell *formula);
int parse_user(const char *path, struct user_data *user_mods);

/// Worksheet/User management
void evaluate_worksheet(struct worksheet *ws); // Maybe add other parameters ?
void apply_user(struct worksheet *ws, struct user_data *user_mods);
void produce_view(struct worksheet *ws, const char *path); // other parameters ?
void produce_changes(struct worksheet *ws, const char *path); // same thing ?
void release_worksheet(struct worksheet *ws);
void release_user(struct user_data *user_mods);

/// Print things
void print_worksheet(struct worksheet *ws);
void print_user(struct user_data *user_mods);
void print_data(struct cell *value, const char *separator);

/// Utilities
enum cell_ty token_type(const char *token);

#endif

