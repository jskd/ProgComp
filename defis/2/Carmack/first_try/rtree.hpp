#ifndef _R_TREE_HPP
#define _R_TREE_HPP

#include <vector>
#include <stack>
#include <functional>

#include "cell.hpp"
#include "parser.hpp"

class INode;

class Leaf;

class Node {
public:
    /* Bounding Box */
    Area<int> bb;
    INode *parent;
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
    const int CAPACITY = 2;

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
    const int CAPACITY = 4;
    std::vector<FormulaNode *> formulas;

    Leaf();
    void split(FormulaNode *f, int *id);
    void update_bb();
public:
    void insert(FormulaNode &f);
    void search(point &p, std::stack<FormulaNode *> &fs);
    void foreach(std::function<void(FormulaNode &)> fun);
};

INode *roots(Parser &p, std::vector<FormulaNode *> &out);

#endif
