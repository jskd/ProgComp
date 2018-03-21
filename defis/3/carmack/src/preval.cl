
typedef struct {
    int x1, x2, y1, y2, v;
    int res;
} formula;

typedef struct {
    int x, y, v;
} value;

__kernel void preval(__global formula *fs, __global value *vs,
                     __local int *one_value_res) {
    int group_size = get_local_size(0);
    int local_id = get_local_id(0);
    int group_id = get_group_id(0);

    int offset, mask;

    if(vs[local_id].v == fs[group_id].v  &&
       vs[local_id].x >= fs[group_id].x1 &&
       vs[local_id].x <= fs[group_id].x2 &&
       vs[local_id].y >= fs[group_id].y1 &&
       vs[local_id].y <= fs[group_id].y2) {
        one_value_res[local_id] = 1;
    }
    else
        one_value_res[local_id] = 0;

    for(offset = 1; offset < group_size; offset *= 2) {
        mask = 2 * offset - 1;
        barrier(CLK_LOCAL_MEM_FENCE);
        if(!(local_id & mask) && local_id + offset < group_size)
            one_value_res[local_id] += one_value_res[local_id + offset];
    }

    barrier(CLK_LOCAL_MEM_FENCE);
    if(local_id == 0)
        fs[group_id].res = one_value_res[0];
}
