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

bool insert_r(r_tree_node *r, node &formula){
	area a = area_of_node(formula);
	r_tree_node *candidate, *old_r;
	while(true){
		if (equal_a(r->a, a)){
			//node with exact boundaries exists - add to its formulae vector
			r->formulae.push_back(&formula);
			return true;
		}
		old_r = r;
		vector<r_tree_node *>::iterator i = r->subtrees.begin();
		while(i != r->subtrees.end()){
			candidate = *i++;
			if(inside_or_equal_a(candidate->a, a)){
				r = candidate;
				break;
			}
		}
		if(old_r == r)
			//need to create new node for this formula here
			return insert_new_r(r, formula);
	}
}

void clever_divider(r_tree_node **nodes, int len, r_tree_node **pair){
	//rewrite, make cleverer
	int pivot = len / 2, i;
	area a;
	fill_a(a, 0,0,0,0);
	for(i = 0; i < pivot; i++){
		pair[0]->subtrees.push_back(nodes[i]);
		a = least_upper_a(a, nodes[i]->a);
	}
	pair[0]->a = a;
	
	fill_a(a, 0,0,0,0);
	for(i = pivot; i < len; i++){
		pair[1]->subtrees.push_back(nodes[i]);
		a = least_upper_a(a, nodes[i]->a);
	}
	pair[1]->a = a;
	return;
}

void adventure_time(vector<r_tree_node *> v, r_tree_node** ret){
	ret[0] = new r_tree_node;
	ret[1] = new r_tree_node;
	int len = v.size(), j = 0;
	vector<r_tree_node *>::iterator i = v.begin();
	r_tree_node *nodes[len];
	while(i != v.end()){
		nodes[j++] = *i++;
	}
	clever_divider(nodes, len, ret);	
	return;
}

bool insert_new_r(r_tree_node *r, node &formula){
	area a = area_of_node(formula);
	r_tree_node *new_node = new r_tree_node; //attention! "new" here!
	new_node->a = a;
	new_node->formulae.push_back(&formula);
	
		r->subtrees.push_back(new_node);
	if(r->subtrees.size() <= MAX_SUBTREES){
		return true;
	}
	r_tree_node *pair[2];
	adventure_time(r->subtrees, pair);
	r->subtrees.clear();
	r->subtrees.push_back(pair[0]);
	r->subtrees.push_back(pair[1]);
	return true;
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