#ifndef _PARSER_HPP
#define _PARSER_HPP

#include <fstream>

enum cell_type { Value, Formula, None };

struct cell {
    enum cell_type type;
    union {
        struct { int value; };
        struct { int x1, y1, x2, y2, v; };
    };
};

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
