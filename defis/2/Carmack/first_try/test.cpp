#include <iostream>
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

    bool next_value(cell *c);
public:
    bool eof = false, eol = false;

    Parser(string data) :
        fs(*(new fstream(data, fstream::in))) { }

    bool next_cell(cell *c);
};

void fill_with_value(cell *c, int v) {
    c->type = Value;
    c->value = v;
}

bool Parser::next_value(cell *c) {
    char i;
    bool b = true;
    int v = 0;

    while(fs >> noskipws >> i) {
        switch(i) {
        case ' ': case ';': case ')': case ',':
            if(b)
                break;
        case '=':
            eol = false;
            eof = fs.eof();
            fill_with_value(c, v);
            return (i == '=');
        case '\n':
            eol = true;
            eof = fs.eof();
            if(b)
                c->type = None;
            else
                fill_with_value(c, v);
            return false;
        default:
            b = false;
            v = v * 10 + ((int)(i - '0'));
        }
    }

    eol = true;
    eof = fs.eof();
    fill_with_value(c, v);
    return false;
}

bool Parser::next_cell(cell *c) {
    char i;
    int x1, y1, x2, y2, v;

    if(!next_value(c))
        return eol;

    fs >> noskipws >> i;
    fs >> noskipws >> i;
    next_value(c);
    x1 = c->value;
    next_value(c);
    y1 = c->value;
    next_value(c);
    x2 = c->value;
    next_value(c);
    y2 = c->value;
    next_value(c);
    v = c->value;

    c->type = Formula;
    c->x1 = x1;
    c->y1 = y1;
    c->x2 = x2;
    c->y2 = y2;
    c->v = v;
    return eol;
}
