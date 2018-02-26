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
    area bb;
    INode *parent;
    virtual void insert(formula &f) = 0;
    virtual void search(point &p, std::stack<formula *> &fs) = 0;
    virtual void foreach(std::function<void(formula &)> fun) = 0;
    void preval(Parser &p);
    Node &operator+=(formula &f);
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
    void insert(formula &f);
    void search(point &p, std::stack<formula *> &fs);
    void foreach(std::function<void(formula &)> fun);
    void addINode(INode *n);
    void addLeaf(Leaf *l);
};

class Leaf : public Node {
    friend class INode;
private:
    const int CAPACITY = 4;
    std::vector<formula *> formulas;

    Leaf();
    void split(formula *f, int *id);
    void update_bb();
public:
    void insert(formula &f);
    void search(point &p, std::stack<formula *> &fs);
    void foreach(std::function<void(formula &)> fun);
};

INode *roots(Parser &p, std::vector<formula *> &out);

#endif
