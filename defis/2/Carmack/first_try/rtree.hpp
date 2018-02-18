#ifndef _R_TREE_HPP
#define _R_TREE_HPP

#include <vector>
#include <stack>
#include <functional>

#include "cell.hpp"
#include "parser.hpp"

using namespace std;

class INode;

class Leaf;

class Node {
public:
    /* Bounding Box */
    area bb;
    INode *parent;
    virtual void insert(formula &f) = 0;
    virtual void search(point &p, stack<formula *> &fs) = 0;
    virtual void foreach(function<void(formula &)> fun) = 0;
    Node &operator+=(formula &f);
};

class INode : public Node {
private:
    static INode *head;
    vector<Node *> childs;
    const int CAPACITY = 2;

    INode();
    INode(INode *child);
    INode *wrap(Leaf *l);
public:
    static INode *getHead();
    void insert(formula &f);
    void search(point &p, stack<formula *> &fs);
    void foreach(function<void(formula &)> fun);
    void addINode(INode *n);
    void addLeaf(Leaf *l);
};

class Leaf : public Node {
    friend class INode;
private:
    const int CAPACITY = 4;
    vector<formula *>formulas;

    Leaf();
    void split(formula *f, int *id);
    void update_bb();
public:
    void insert(formula &f);
    void search(point &p, stack<formula *> &fs);
    void foreach(function<void(formula &)> fun);
};

INode *roots(Parser &p, vector<formula *> &out);

#endif
