#ifndef _R_TREE_HPP
#define _R_TREE_HPP

#include <vector>
#include <stack>
#include <functional>

#include "formulanode.hpp"
#include "parser.hpp"

////////////////////////////////////////////////////////////////
///
///  This is a direct R-tree implementation, if you have some
///  questions, it means that either:
///      1. You don't know something about r-trees, or
///      2. There is an error in this implementation.
///  In any of this two cases, it's better to pretend that you
///  have no questions.
///
///  This structure contains all formulas read from file,
///  it helps accelerate formula search by the given position.
///
////////////////////////////////////////////////////////////////

/*
  There are two types of node: Internal Node (INode) and Leaf.
  Internal Node can contain any other node, while Leaf
  contains formulas. */
class INode;

class Leaf;

class Node {
public:
    Area bb;
    INode *parent = NULL;
    virtual void insert(FormulaNode &f) = 0;
    virtual void search(point &p, std::stack<FormulaNode *> &fs) = 0;
    virtual void foreach(std::function<void(FormulaNode &)> fun) = 0;
    void preval(Parser &p);
    Node &operator+=(FormulaNode &f);
};

class INode : public Node {
private:
    static INode *head;
    std::vector<Node *> children;
    const size_t CAPACITY = 20;

    INode();
    INode(INode *child);
    INode *wrap(Leaf *l);
public:
    static INode *getHead();
    void insert(FormulaNode &f);
    void search(point &p, std::stack<FormulaNode *> &fs);
    void foreach(std::function<void(FormulaNode &)> fun);
    void addINode(INode *n);
    void addLeaf(Leaf *l);
};

class Leaf : public Node {
    friend class INode;
private:
    const size_t CAPACITY = 40;
    std::vector<FormulaNode *> formulas;

    Leaf();
    void split(FormulaNode *f, int *id);
    void update_bb();
public:
    void insert(FormulaNode &f);
    void search(point &p, std::stack<FormulaNode *> &fs);
    void foreach(std::function<void(FormulaNode &)> fun);
};

////////////////////////////////////////////////////////////////
///
///  [roots(p, out)] return head of the given r-tree, constructed
///  from [p], and fill [out] vector with independant formulas
///  (formulas of the level 0).
///
////////////////////////////////////////////////////////////////
INode *roots(Parser &p, std::vector<FormulaNode *> &out);

#endif
