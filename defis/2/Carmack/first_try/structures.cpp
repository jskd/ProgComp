#include "structures.hpp"

void fill_a(area &a, unsigned x1, unsigned y1, unsigned x2, unsigned y2){
	#define FILL(s) a.s = s
	FILL(x1);
	FILL(x2);
	FILL(y1);
	FILL(y2);
	#undef FILL
}

bool is_in_a(area a, unsigned x, unsigned y){
	return x >= a.x1 && x <= a.x2 && y >= a.y1 && y <= a.y2;
} 

bool equal_a(area a, area a2){
	#define EQ(s) a.s == a2.s
	return EQ(x1) && EQ(x2) && EQ(y1) && EQ(y2);
	#undef EQ
}

bool inside_or_equal_a(area outer, area inner){
	return outer.x1 <= inner.x1 && outer.y1 <= inner.y1 &&
		outer.x2 >= inner.x2 && outer.y2 >= inner.y2;
}

area area_of_node(node n){
	area ret;
	fill_a(ret, n.f.x1, n.f.y1, n.f.x2, n.f.y2);
	return ret;
}

area least_upper_a(area a, area a2){
	area ret;
	#define MIN(a, b) a < b ? a : b
	#define MAX(a, b) a > b ? a : b
	fill_a(ret, MIN(a.x1, a2.x1), MIN(a.y1, a2.y1),
		MAX(a.x2, a2.x2), MAX(a.y2, a2.y2));
	return ret;
	#undef MIN
	#undef MAX
}