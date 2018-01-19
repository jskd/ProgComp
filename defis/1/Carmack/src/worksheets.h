#ifndef WORKSHEETS_H
#define WORKSHEETS_H

// Data types

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

// Functions

int parse_data(const char* path, worksheet* ws);
int parse_formula(char* token, data* formula);

void release_worksheet(worksheet* ws);

void print_worksheet(worksheet* ws);

data_ty token_type(const char* token);

#endif
