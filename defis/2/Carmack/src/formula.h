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

/*
  Formula structure is encoded into cl_int8
  (OpenCl vector type) */

typedef cl_int8 formula;

#define X1(f)    (f)->s[0]
#define X2(f)    (f)->s[1]
#define Y1(f)    (f)->s[2]
#define Y2(f)    (f)->s[3]
#define V(f)     (f)->s[4]
#define R(f)     (f)->s[5]
#define LEVEL(f) (f)->s[6]

#define VALUE_MAX_CAPACITY 256

////////////////////////////////////////////////////////////////
///
///  [new_formula(...)] register some formula. Basically, it just
///  allocate some memory in some internal buffer.
///
////////////////////////////////////////////////////////////////
int new_formula(int x1, int y1, int x2, int y2, int v);

formula *get_formula(int id);

void init_opencl();

////////////////////////////////////////////////////////////////
///
///  [push_value(...)] register some value. Basically, it just
///  allocate some memory in some internal buffer.
///  Return non-zero value if number of registered values equals
///  to VALUE_MAX_CAPACITY. Buffer overwrites when trying to push
///  to full buffer.
///
////////////////////////////////////////////////////////////////
int push_value(int x, int y, int v);

////////////////////////////////////////////////////////////////
///
///  [launch_calculation(...)] launch partianl formula calculation
///  over registered formulas and values.
///
////////////////////////////////////////////////////////////////
void launch_calculation();

#endif
