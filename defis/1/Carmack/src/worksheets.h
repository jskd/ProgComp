#ifndef WORKSHEETS_H
#define WORKSHEETS_H

// Data types

// Some data types may be unnecessary for the moment
// but keep in mind that it is a naive implementation.
// Refactor if you want to but please do not loose semantic value !
// Besides, if we implement a big data worksheet engine in the future,
// we may need them for storing multiple metadata.

enum data_ty { INVALID, FORMULA, VALUE };

struct formula {
	int r1, r2, c1, c2, val;
};

struct data {
	enum data_ty ty;
	union {
		int value;
		struct formula formula;
	} udata;
};

struct line {
	unsigned elements_count;
	struct data *content;
};

struct worksheet {
	unsigned lines_count;
	struct line *lines;
};

struct user_component {
	int r, c;
	struct data value;
};

struct user {
	unsigned changes_count;
	struct user_component *content;
};

// Functions

/// Parsing
int parse_data(const char *path, struct worksheet *ws);
int parse_formula(char *token, struct data *formula);
int parse_user(const char *path, struct user *user_mods);

/// Worksheet/User management
void evaluate_worksheet(struct worksheet *ws); // Maybe add other parameters ?
void apply_user(struct worksheet *ws, struct user *user_mods);
void produce_view(struct worksheet *ws, const char *path); // add other parameters ?
void produce_changes(struct worksheet *ws, const char *path); // same thing ?
void release_worksheet(struct worksheet *ws);
void release_user(struct user *user_mods);

/// Print things
void print_worksheet(struct worksheet *ws);
void print_user(struct user *user_mods);
void print_data(struct data *value, const char *separator);

/// Utilities
enum data_ty token_type(const char *token);

#endif
