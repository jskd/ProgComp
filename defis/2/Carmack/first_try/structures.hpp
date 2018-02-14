#ifndef _STRUCTURES_HPP
#define _STRUCTURES_HPP

#define MAX_SUBTREES 3

#include <vector> 

using namespace std;

enum cell_type { Value, Formula, None };

typedef struct {
    enum cell_type type;
    union {
        struct { int value; };
        struct { int x1, y1, x2, y2, v; };
    };
} cell;

typedef struct _node{
	cell f;
	int pos_x, pos_y;
	vector<_node> children, parents;
} node;

typedef vector<node *> graph;

typedef struct{
	unsigned x1, y1, x2, y2;
}area;


typedef struct rtn{
	area a; 
	vector<node *> formulae;
	vector<rtn *> subtrees;
} r_tree_node;

void fill_a(area &a, unsigned x1, unsigned y1, unsigned x2, unsigned y2);
bool is_in_a(area a, unsigned x, unsigned y); //bounds included
bool equal_a(area a, area a2);
bool inside_or_equal_a(area outer, area inner); 
area least_upper_a(area a, area a2);
area area_of_node(node n);

bool insert_r(r_tree_node *r, node &formula); 
bool insert_new_r(r_tree_node *r, node &formula); 

#endif
