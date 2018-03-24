#ifndef _CELL_HPP
#define _CELL_HPP

#include <vector>
#include <ostream>

#include "area.hpp"
#include "parser.hpp"

extern "C" {
    #include "formula.h"
}

////////////////////////////////////////////////////////////////
///
///  If formula structure (from "formula.h") is just a low level
///  formula representation, especially useful for OpenCL,
///  FormulaNode in contrast is more like a single node of some
///  dependence graph.
///
////////////////////////////////////////////////////////////////

class FormulaNode {
private:
    int formula_id;
public:
    point p;
    std::vector<FormulaNode *> children;
    std::vector<FormulaNode *> parents;

    FormulaNode(int x1, int y1, int x2, int y2, int v);
    FormulaNode(cell &c);

    /*
      As everywhere in this project, 'bb' stands for bounding box.
      In this context it's just a formula area. */
    Area bb();

    int value();

    void value(int v);

    int result();

    void result(int r);

    int level();

    void level(int l);

    friend std::ostream &operator<<(std::ostream &out, FormulaNode &f);
};

////////////////////////////////////////////////////////////////
///
///  [normalize(roots, top)] fill [top] vector with all formulas
///  accessible from at least one element of [roots] in a
///  topological order. So, strictly speaking, if [roots] is a
///  collection of independent formulas (level = 0),
///  then we have this property:
///  top[i]->level < top[j]->level iff i < j.
///
////////////////////////////////////////////////////////////////
void normalize(std::vector<FormulaNode *> &roots,
               std::vector<FormulaNode *> &top);

////////////////////////////////////////////////////////////////
///
///  [evaluate(top)] evaluate formulas listed in [top] vector.
///  It's silently assumed that [top] sorted in topological order.
///
////////////////////////////////////////////////////////////////
void evaluate(std::vector<FormulaNode *> &top);

////////////////////////////////////////////////////////////////
///
///  [superior_preval(p)] pre-evaluation of all registered
///  formulas (see "formula.h" for more information about
///  "formula registration"). In contrast to [preval] of rtree,
///  this function use GPU (over OpenCL), and in contrast to
///  [preval] of "defi 3" this function doesn't work.
///
////////////////////////////////////////////////////////////////
void superior_preval(Parser &p);

int write_view(Parser &p, std::string view);

#endif
