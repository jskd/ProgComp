#ifndef _FORMULA_H
#define _FORMULA_H

// Yes, Luxon, it's a C!
// You're welcome to rewrite this for C++.

#define CL_USE_DEPRECATED_OPENCL_1_2_APIS
#ifdef __APPLE__
    #include <OpenCL/cl.h>
#else
    #include <CL/cl.h>
#endif

// Yes, I tried hard to satisfy OpenCL,
// but he's a tough guy.
typedef cl_int8 formula;

#define X1(f)    (f)->s0
#define X2(f)    (f)->s1
#define Y1(f)    (f)->s2
#define Y2(f)    (f)->s3
#define V(f)     (f)->s4
#define R(f)     (f)->s5
#define LEVEL(f) (f)->s6

int new_formula(int x1, int y1, int x2, int y2, int v);

formula *get_formula(int id);

void init_opencl();

int push_value(int x, int y, int v);

void launch_calculation();

#endif
