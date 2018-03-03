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
    int formula_id;
public:
    point p;
    std::vector<FormulaNode *> children;
    std::vector<FormulaNode *> parents;

    FormulaNode(int x1, int y1, int x2, int y2, int v);
    FormulaNode(cell &c);

    Area bb();

    int value();

    void value(int v);

    int result();

    void result(int r);

    int level();

    void level(int l);

    friend std::ostream &operator<<(std::ostream &out, FormulaNode &f);
};

void normalize(std::vector<FormulaNode *> &roots,
               std::vector<FormulaNode *> &top);

void evaluate(std::vector<FormulaNode *> &top);

void superior_preval();

#endif
