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

int new_formula(int x1, int y1, int x2, int y2, int v) {
    formula *f;

    if(!fs || f_size >= f_capacity) {
        f_capacity += STEP;
        fs = realloc(fs, f_capacity * sizeof *fs);
        if(!fs)
            return -1;
    }
    f = fs + f_size;

    f->x1 = x1;
    f->x2 = x2;
    f->y1 = y1;
    f->y2 = y2;
    f->v = v;
    f->r = 0;
    f->level = -1;

    return f_size++;
}

formula *get_formula(int id) {
    return fs + id;
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

    rewind(f);
    size = fread(buf, 1, size, f);
    buf[size] = 0;
    fclose(f);
    return buf;

 exit_with_error:
    fclose(f);
    return NULL;
}

typedef struct __attribute__((aligned)) {
    int x, y, v;
} value;

#define VALUE_MAX_CAPACITY 256

value *vs;
size_t v_size;

cl_platform_id   platform;
cl_device_id     device;
cl_context       context;
cl_program       program;
cl_kernel        kernel;
cl_command_queue queue;

#define CHECK_ERROR(errmsg)                      \
    if(error != CL_SUCCESS ) {                   \
        fprintf(stderr, "%s\n", errmsg);         \
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
        fprintf(stderr, "Can't locate or load OpenCL source file.\n");
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

    vs = malloc(VALUE_MAX_CAPACITY * sizeof *vs);
    if(!vs) {
        fprintf(stderr, "Not enough memory space for 'values' allocation.\n");
        exit(1);
    }
}

int push_value(int x, int y, int v) {
    if(v_size == VALUE_MAX_CAPACITY)
        v_size = 0;

    vs[v_size].x = x;
    vs[v_size].y = y;
    vs[v_size].v = v;

    return (v_size++) == VALUE_MAX_CAPACITY;
}

void launch_calculation() {
    cl_int error;
    cl_mem f_buffer, v_buffer;
    size_t global_work_size[] = { f_size * v_size };
    size_t local_work_size[]  = { v_size };
    cl_event event;

    f_buffer = clCreateBuffer(context, CL_MEM_READ_WRITE | CL_MEM_COPY_HOST_PTR,
                              f_size * sizeof *fs, fs, &error);
    CHECK_ERROR("Unable to create formulas buffer object.");

    v_buffer = clCreateBuffer(context, CL_MEM_READ_ONLY | CL_MEM_COPY_HOST_PTR,
                              v_size * sizeof *vs, vs, &error);
    CHECK_ERROR("Unable to create values buffer object.");

    error = clSetKernelArg(kernel, 0, sizeof(cl_mem), &f_buffer);
    CHECK_ERROR("Unable to initialize first argument.");
    error = clSetKernelArg(kernel, 1, sizeof(cl_mem), &v_buffer);
    CHECK_ERROR("Unable to initialize second argument.");

    error = clEnqueueNDRangeKernel(queue, kernel, 1, NULL,
                                   global_work_size, local_work_size,
                                   0, NULL, &event);
    CHECK_ERROR("Unable to enqueue to command-queue.");
    clWaitForEvents(1, &event);

    error = clEnqueueReadBuffer(queue, f_buffer, CL_TRUE, 0,
                                f_size * sizeof *fs, fs, 0, NULL, NULL);
    CHECK_ERROR("Unable to enqueue to buffer.");
}
