#ifndef _PARSER_HPP
#define _PARSER_HPP

#include "structures.hpp"
#include <fstream>

class Parser {
private:
    fstream &fs;
    bool _eof = false, _eol = false;

    bool next_value(cell *c);
public:
    const bool &eof;
    const bool &eol;

    Parser(string data);

    bool next_cell(cell *c);
};

#endif
