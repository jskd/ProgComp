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

class formula {
public:
    area bb;
    int value;
    point p;

    // Should be private
    int level = -1;
    int result = 0;
    std::vector<formula *> children;
    std::vector<formula *> parents;

    formula();
    formula(int x1, int y1, int x2, int y2, int v);
    formula(cell &c);

    friend std::ostream &operator<<(std::ostream &out, const formula &f);
};

void normalize(std::vector<formula *> &roots, std::vector<formula *> &top);

void evaluate(std::vector<formula *> &top);

#endif
