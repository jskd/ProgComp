#include "first_pass.hpp"
#include "parser.hpp"

r_tree_node *get_root(){
	r_tree_node *ret = new r_tree_node;
	fill_a(ret->a, 0,0, UINT_MAX, UINT_MAX);
	return ret;
}

void clever_divider(r_tree_node **nodes, int len, r_tree_node **pair){
	//rewrite, make cleverer
	int pivot = len / 2, i;
	area a;
	fill_a(a, UINT_MAX, UINT_MAX, 0, 0);
	for(i = 0; i < pivot; i++){
		pair[0]->subtrees.push_back(nodes[i]);
		a = least_upper_a(a, nodes[i]->a);
	}
	pair[0]->a = a;
	
	fill_a(a, UINT_MAX, UINT_MAX, 0, 0);
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

bool insert_new_r(r_tree_node *r, node *formula){
	area a = area_of_node(*formula);
	r_tree_node *new_node = new r_tree_node; //attention! "new" here!
	new_node->a = a;
	new_node->formulae.push_back(formula);
	
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

bool insert_r(r_tree_node *r, node *formula){
	area a = area_of_node(*formula);
	r_tree_node *candidate, *old_r;
	while(true){
		if (equal_a(r->a, a)){
			//node with exact boundaries exists - add to its formulae vector
			r->formulae.push_back(formula);
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

bool add_child(node *child, node *parent){
	parent->children.push_back(child);
	child->parents.push_back(parent);
}
bool add_child(node *formula, r_tree_node *r){
	if (is_in_a(r->a, formula->pos_x, formula->pos_y)){
		vector<node *>::iterator i;
		for(i = r->formulae.begin(); i != r->formulae.end(); i++)
			add_child(formula, *i);
		vector<r_tree_node *>::iterator j;
		for(j = r->subtrees.begin(); j != r->subtrees.end(); j++)
			if(is_in_a((*j)->a, formula->pos_x, formula->pos_y))
				add_child(formula, *j);
		return true;
	}
	return false;
}

bool insert_f(node *formula, f_hash *hash){
	long long unsigned key = point_to_key(formula->pos_x, formula->pos_y);
	hash->insert(pair<long long unsigned, node *>(key, formula));
}

bool search_f(unsigned x, unsigned y, node **result, f_hash *hash){
	std::unordered_map<long long unsigned, node *>::const_iterator found;
	found = hash->find(point_to_key(x, y));
	if(found == hash->end())
		return false;
	*result = found->second;
	return true;
}

bool adjust_children(f_hash *hash, node *parent,
	unsigned current_x, unsigned current_y){
	unsigned x = parent->f.x1, y = parent->f.y1;
	node *child;
	while(y <= parent->f.y2 && y <= current_y){
		if(x >= current_x && y == current_y)
			break; //don't add formula to its children the second time, please
		if(search_f(x, y, &child, hash))
			add_child(child, parent);
		x++;
		if(x > parent->f.x2){
			x = parent->f.x1;
			y++;
		}
	}
	return true;
}

bool first_pass(Parser &p, f_hash *hash, r_tree_node *r_tree, graph *g){
	unsigned x = 0, y = 0;
	r_tree = get_root();
	g = new graph;
	hash = new f_hash;
	cell c;
	node *n;
	while(true){
		p.next_cell(&c);
		if(c.type == Formula){
			n = new node;
			n->f = c;
			n->pos_x = x;
			n->pos_y = y;
			insert_r(r_tree, n);
			add_child(n, r_tree);
			adjust_children(hash, n, x, y);
		}
		
		x++;
		if(p.eol){
			x = 0;
			y++;
		}
		if(p.eof)
			break;
	}
	return true;
}