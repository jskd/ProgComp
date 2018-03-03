#include <stdlib.h>
#include <unistd.h>
#include <stdio.h>
#define CL_USE_DEPRECATED_OPENCL_1_2_APIS
#include <CL/cl.h>

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

char *load_cl_source() {
    char *buf;
    long size;
    FILE *f = fopen("preval.cl", "r");

    if(!f)
        return NULL;

    if(fseek(f, 0, SEEK_END))
        goto exit_with_error;

    size = ftell(f);
    if(size == -1)
        goto exit_with_error;

    buf = malloc(size + 1);
    if(!buf)
        goto exit_with_error;

    size = fread(buf, 1, size, f);
    buf[size] = 0;
    return buf;

 exit_with_error:
    fclose(f);
    return NULL;
}

cl_platform_id   platform;
cl_device_id     device;
cl_context       context;
cl_program       program;
cl_kernel        kernel;
cl_command_queue queue;

#define CHECK_ERROR(errmsg)                      \
    if(error != CL_SUCCESS ) {                   \
        perror(errmsg);                          \
        exit(1);                                 \
    }

void init_opencl() {
    char *source;
    cl_int error;

    error = clGetPlatformIDs(1, &platform, NULL);
    CHECK_ERROR("Unable to find any OpenCL platforms.");

    error = clGetDeviceIDs(platform, CL_DEVICE_TYPE_GPU, 1, &device, NULL);
    CHECK_ERROR("Can't locate any OpenCL compliant GPU device.");

    context = clCreateContext(NULL, 1, &device, NULL, NULL, &error);
    CHECK_ERROR("Can't create a valid OpenCL context.");

    source = load_cl_source();
    if(!source) {
        perror("Can't locate or load OpenCL source file.");
        exit(1);
    }
    program = clCreateProgramWithSource(context, 1, (const char **)&source,
                                        NULL, &error);
    free(source);
    CHECK_ERROR("Can't create the OpenCL program object.");

    error = clBuildProgram(program, 1, &device, NULL, NULL, NULL);
    CHECK_ERROR("Can't create OpenCL program, possibly faulty source file.");

    error = clCreateKernelsInProgram(program, 1, &kernel, NULL);
    CHECK_ERROR("Unable to create kernel object.");

    queue = clCreateCommandQueue(context, device, 0, &error);
    CHECK_ERROR("Unable to create command-queue.");
}
