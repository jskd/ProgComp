#include "first_pass.hpp"



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
