#ifndef _FORMULA_H
#define _FORMULA_H

#define CL_USE_DEPRECATED_OPENCL_1_2_APIS
#ifdef __APPLE__
    #include <OpenCL/cl.h>
#else
    #include <CL/cl.h>
#endif

typedef struct {
    int x1, x2, y1, y2, v;
    int res;
} formula;

int new_formula(int x1, int y1, int x2, int y2, int v);

formula *get_formula(int id);

void init_opencl();

int push_value(int x, int y, int v);

void launch_calculation();

#endif
