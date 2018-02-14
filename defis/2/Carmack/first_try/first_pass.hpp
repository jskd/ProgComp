#ifndef FIRST_PASS_HPP
#define FIRST_PASS_HPP

#include <vector>
#include "parser.hpp"

using namespace std;

int pass_file(&Parser file_p, &vector<unsigned> first_lines);

#endif
