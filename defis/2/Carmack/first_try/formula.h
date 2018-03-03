#ifndef _FORMULA_H
#define _FORMULA_H

typedef struct __attribute__((aligned)) {
    int x1, x2, y1, y2, v, r;
} formula;

formula *new_formula(int x1, int y1, int x2, int y2, int v);

void init_opencl();

#endif
