#include "parser.hpp"

using namespace std;

Parser::Parser(string data) :
    fs(*(new fstream(data, fstream::in))),
    eof(_eof),
    eol(_eol) { }

Parser::~Parser() {
    fs.close();
}

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
            _eol = false;
            _eof = fs.eof();
            fill_with_value(c, v);
            return (i == '=');
        case '\n':
            _eol = true;
            _eof = fs.eof();
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

    _eol = true;
    _eof = fs.eof();
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

#define FILL_FIELD(f) \
    next_value(c);    \
    f = c->value

    FILL_FIELD(x1);
    FILL_FIELD(y1);
    FILL_FIELD(x2);
    FILL_FIELD(y2);
    FILL_FIELD(v);

#undef  FILL_FIELD
#define FILL_FIELD(f) \
    c->f = f

    c->type = Formula;
    FILL_FIELD(x1);
    FILL_FIELD(y1);
    FILL_FIELD(x2);
    FILL_FIELD(y2);
    FILL_FIELD(v);
    return eol;
}

void Parser::reposition() {
    fs.seekg(0);
}
