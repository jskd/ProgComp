#ifndef WORKSHEETS_H
#define WORKSHEETS_H

// Data types

// Some data types may be unnecessary for the moment
// but keep in mind that it is a naive implementation.
// Refactor if you want to but please do not loose semantic value !
// Besides, if we implement a big data worksheet engine in the future,
// we may need them for storing multiple metadata.

typedef enum { INVALID, FORMULA, VALUE } data_ty;

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

typedef struct {
	unsigned elements_count;
	data *content;
} line;

typedef struct {
	unsigned lines_count;
	line* lines;
} worksheet;

typedef struct {
	int r, c;
	data value;
} user_component;

typedef struct {
	unsigned changes_count;
	user_component* content;
} user;

// Functions

/// Parsing
int parse_data(const char* path, worksheet* ws);
int parse_formula(char* token, data* formula);
int parse_user(const char* path, user* user_mods);

/// Worksheet/User management
void evaluate_worksheet(worksheet* ws); // Maybe add others parameters ?
void apply_user(worksheet* ws);
void release_worksheet(worksheet* ws);
void release_user(user* user_mods);

/// Print things
void print_worksheet(worksheet* ws);
void print_user(user* user_mods);
void print_data(data* value, const char* separator);

/// Utilities
data_ty token_type(const char* token);

#endif
