#include <cassert>
#include <cstdlib>

#include "rtree.hpp"

Node &Node::operator+=(formula &f) {
    insert(f);
    return *this;
}

INode::INode() {
    parent = NULL;
}

INode::INode(INode *child) {
    child->parent = this;
    bb = child->bb;
    children.push_back(child);
}

void INode::insert(formula &f) {
    vector<Node *>::iterator i;
    Node *n;
    area a;
    int s, old_s = -1;

    if(!children.size()) {
        Leaf *l = new Leaf();
        l->bb = f.bb;
        addLeaf(l);
        insert(f);
        return;
    }

    for(i = children.begin(); i != children.end(); ++i) {
        if(contains((*i)->bb, f.bb)) {
            (**i) += f;
            return;
        }
        intersect((*i)->bb, f.bb, a);
        s = surface(a);
        if(old_s < 0 || s < old_s) {
            old_s = s;
            n = *i;
        }
    }
    intersect(n->bb, f.bb, n->bb);
    *n += f;
}

void INode::search(point &p, stack<formula *> &fs) {
    vector<Node *>::iterator i;

    for(i = children.begin(); i != children.end(); ++i) {
        if(pcontains((*i)->bb, p))
            (*i)->search(p, fs);
    }
}

void INode::foreach(function<void(formula &)> fun) {
    vector<Node *>::iterator i;

    for(i = children.begin(); i != children.end(); ++i)
        (*i)->foreach(fun);
}

INode *INode::wrap(Leaf *l) {
    INode *n = new INode();
    n->addLeaf(l);
    n->bb = l->bb;
    return n;
}

void INode::addINode(INode *n) {
    if(children.size() < CAPACITY) {
        children.push_back(n);
        n->parent = this;
        return;
    }

    if(!parent) {
        new INode(this);
        if(this == head)
            head = parent;
    }
    parent->addINode(n);
}

void INode::addLeaf(Leaf *l) {
    if(children.size() < CAPACITY) {
        children.push_back(l);
        l->parent = this;
        return;
    }

    if(!parent) {
        new INode(this);
        if(this == head)
            head = parent;
    }
    parent->addINode(wrap(l));
}

INode *INode::head = new INode();

INode *INode::getHead() {
    return head;
}

Leaf::Leaf() { }

void Leaf::insert(formula &f) {
    int id;
    formula *f2;
    Leaf *l;

    if(formulas.size() < CAPACITY) {
        intersect(bb, f.bb, bb);
        formulas.push_back(&f);
        return;
    }
    split(&f, &id);
    f2 = formulas[id];
    formulas[id] = formulas[CAPACITY - 1];
    formulas.pop_back();
    update_bb();
    l = new Leaf();
    *l += f;
    *l += *f2;
    parent->addLeaf(l);
}

void Leaf::search(point &p, stack<formula *> &fs) {
    vector<formula *>::iterator i;

    for(i = formulas.begin(); i != formulas.end(); ++i) {
        if(pcontains((*i)->bb, p))
            fs.push(*i);
    }
}

void Leaf::foreach(function<void(formula &)> fun) {
    vector<formula *>::iterator i;

    for(i = formulas.begin(); i != formulas.end(); ++i)
        fun(**i);
}

void Leaf::split(formula *f, int *id) {
    int i, median = surface(bb) / 2, s, old_s = -1;
    area a;

    for(i = 0; i < formulas.size(); i++) {
        intersect(formulas[i]->bb, f->bb, a);
        s = surface(a);
        if(old_s < 0 || abs(median - s) < old_s) {
            old_s = s;
            *id = i;
        }
    }
}

void Leaf::update_bb() {
    int i;

    assert(formulas.size() > 0);
    bb = formulas[0]->bb;
    for(i = 1; i < formulas.size(); i++)
        intersect(bb, formulas[i]->bb, bb);
}

INode *roots(Parser &p, vector<formula *> &out) {
    cell c;
    formula *f;
    int x = 0, y = 0;
    stack<formula *> fs;
    INode *head = INode::getHead();

    while(!p.eof) {
        p.next_cell(&c);
        switch(c.type) {
        case None: case Value: break;
        case Formula:
            f = new formula(c);
            f->p.x = x;
            f->p.y = y;
            (*head) += (*f);
            head = INode::getHead();
        }

        if(p.eol) {
            y = 0;
            x++;
        }
        else
            y++;
    }

    head->foreach(
        [&fs, &head](formula &formula) {
			head->search(formula.p, fs);
            while(!fs.empty()) {
                formula.children.push_back(fs.top());
                fs.top()->has_parent = true;
                fs.pop();
            }
        }
    );

    head->foreach(
        [&out](formula &formula) {
            if(!formula.has_parent)
                out.push_back(&formula);
        }
    );

    return head;
}
