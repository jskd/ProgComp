#include <stdlib.h>

#include "formula.h"

#define STEP 10

formula *fs;
size_t f_size;
size_t f_capacity;

formula *new_formula(int x1, int y1, int x2, int y2, int v) {
    formula *f;

    if(!fs || f_size == f_capacity) {
        f_capacity += STEP;
        fs = realloc(fs, f_capacity * sizeof *fs);
    }
    f = fs + (f_size++);

    f->x1 = x1;
    f->x2 = x2;
    f->y1 = y1;
    f->y2 = y2;
    f->v  = v;
    f->r  = 0;

    return f;
}
