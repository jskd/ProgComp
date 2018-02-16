#ifndef FIRST_PASS_HPP
#define FIRST_PASS_HPP

#include <climits>
#include "structures.hpp"
#include "parser.hpp"

bool insert_r(r_tree_node *r, node *formula); 
bool insert_new_r(r_tree_node *r, node *formula); 
r_tree_node *get_root_r();
bool insert_f(unsigned x, unsigned y, node *formula, f_hash *hash);
void search_f(unsigned x, unsigned y, node *formula, f_hash *hash);
bool first_pass(Parser &p, f_hash *hash, r_tree_node *r_tree, graph *g);
bool add_child(node *child, node *parent);
bool add_child(node *formula, r_tree_node *r);


#endif
