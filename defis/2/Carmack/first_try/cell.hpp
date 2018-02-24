#ifndef _CELL_HPP
#define _CELL_HPP

#include <vector>
#include <ostream>

#include "area.hpp"

enum cell_type { Value, Formula, None };

struct cell {
    enum cell_type type;
    union {
        struct { int value; };
        struct { int x1, y1, x2, y2, v; };
    };
};

using namespace std;

class formula {
public:
    area bb;
    int value;
    point p;

    // Should be private
    int level = -1;
    int result = 0;
    vector<formula *> children;
    vector<formula *> parents;

    formula();
    formula(int x1, int y1, int x2, int y2, int v);
    formula(cell &c);

    friend ostream &operator<<(ostream &out, const formula &f);
};

void normalize(vector<formula *> &roots, vector<formula *> &top);

void evaluate(vector<formula *> &top);

#endif
