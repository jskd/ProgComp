#ifndef FIRST_PASS_HPP
#define FIRST_PASS_HPP

#include <climits>
#include "structures.hpp"
#include "parser.hpp"
bool insert_r(r_tree_node *r, node &formula); 
bool insert_new_r(r_tree_node *r, node &formula); 
r_tree_node *get_root_r();

bool first_pass(Parser &p, f_hash *hash, r_tree_node *r_tree, graph *g);

#endif
