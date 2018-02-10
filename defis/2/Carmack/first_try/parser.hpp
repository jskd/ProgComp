#ifndef _PARSER_HPP
#define _PARSER_HPP

#include <fstream>

using namespace std;

enum cell_type { Value, Formula, None };

typedef struct {
    enum cell_type type;
    union {
        struct { int value; };
        struct { int x1, y1, x2, y2, v; };
    };
} cell;

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
