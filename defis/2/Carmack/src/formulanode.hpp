#ifndef _CELL_HPP
#define _CELL_HPP

#include <vector>
#include <ostream>

#include "area.hpp"
#include "parser.hpp"

extern "C" {
    #include "formula.h"
}

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

void superior_preval(Parser &p);

int write_view(Parser &p, std::string view);

#endif
