#ifndef _PARSER_HPP
#define _PARSER_HPP

#include <fstream>

#include "cell.hpp"

class Parser {
private:
    std::fstream fs;
    bool _eof = false, _eol = false;

    bool next_value(cell *c);
public:
    const bool &eof;
    const bool &eol;

    Parser(std::string data);
    ~Parser();

    bool next_cell(cell *c);
    void reposition();
};

#endif
