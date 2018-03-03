
typedef struct __attribute__((aligned)) {
    int x1, x2, y1, y2, v, r;
} formula;

typedef struct __attribute__((aligned)) {
    int x, x, v;
} value;

__kernel void preval(__global formula *fs, __global value *vs) {
    int f_id = get_global_id(0);
    int v_id = get_global_id(1);

    if(vs[v_id].x >= fs[f_id].x1 && vs[v_id].x <= fs[f_id].x2 &&
       vs[v_id].y >= fs[f_id].y1 && vs[v_id].y <= fs[f_id].y2 &&
       vs[v_id].v == fs[f_id].v)
        fs[f_id].r++;
}
