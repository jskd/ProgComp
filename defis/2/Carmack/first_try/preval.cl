#pragma OPENCL EXTENSION cl_khr_local_int32_extended_atomics : enable

__kernel void preval(__global int8 *fs, __global int4 *vs, __global int *rs) {
    int f_id = get_group_id(0);
    int v_id = get_global_id(1);

    // If even this doesn't work there is no reason to continue.
    atomic_inc(rs);

    /*
    if(fs[f_id].s6 != -1 &&
       vs[v_id].x >= fs[f_id].s0 && vs[v_id].x <= fs[f_id].s1 &&
       vs[v_id].y >= fs[f_id].s2 && vs[v_id].y <= fs[f_id].s3 &&
       vs[v_id].z == fs[f_id].s4)
        atomic_inc(rs + f_id, 1);
    */
}
