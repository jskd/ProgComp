#ifndef _CELL_HPP
#define _CELL_HPP

#include <vector>
#include <ostream>

#include "area.hpp"

extern "C" {
    #include "formula.h"
}

enum cell_type { Value, Formula, None };

struct cell {
    enum cell_type type;
    union {
        struct { int value; };
        struct { int x1, y1, x2, y2, v; };
    };
};

class FormulaNode {
private:
    formula *f;
public:
    Area<int &> bb;
    int &value;
    point p;
    int &level;
    int &result;
    std::vector<FormulaNode *> children;
    std::vector<FormulaNode *> parents;

    FormulaNode(int x1, int y1, int x2, int y2, int v);
    FormulaNode(cell &c);

    friend std::ostream &operator<<(std::ostream &out, const FormulaNode &f);
};

void normalize(std::vector<FormulaNode *> &roots,
               std::vector<FormulaNode *> &top);

void evaluate(std::vector<FormulaNode *> &top);

#endif
