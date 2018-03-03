#include <queue>

#include "cell.hpp"

using namespace std;

FormulaNode::FormulaNode(int x1, int y1, int x2, int y2, int v) :
    f(new_formula(x1, y1, x2, y2, v)),
    bb(Area<int &>(f->x1, f->y1, f->x2, f->y2)),
    value(f->v), result(f->r), level(f->level) { }

FormulaNode::FormulaNode(cell &c) :
    FormulaNode(c.x1, c.y1, c.x2, c.y2, c.v) { }

ostream &operator<<(ostream &out, const FormulaNode &f) {
    out << "# Formula #\n" << "\tposition: x = " << f.p.x << ", y = " <<
        f.p.y << "\n\tx1 = " << f.bb.x1 << ", y1 = " << f.bb.y1 <<
        ", x2 = " << f.bb.x2 << ", y2 = " << f.bb.y2 << ", v = " <<
        f.value << "\n\tactual result: " << f.result << "\n\tlevel = " <<
        f.level << endl;
}

void normalize(vector<FormulaNode *> &roots, vector<FormulaNode *> &top) {
    queue<FormulaNode *> q;
    bool b;

    for(FormulaNode *root : roots) {
        for(FormulaNode *child : root->children) {
            child->level = 1;
            q.push(child);
        }
        // Zero-level formulae are already calculated,
        // so there is no reason to add them to top.
    }

    while(!q.empty()) {
        for(FormulaNode *child : q.front()->children) {
            if(child->level != -1)
                continue;
            b = true;
            for(FormulaNode *parent : child->parents) {
                if(parent->level == -1) {
                    b = false;
                    break;
                }
            }
            if(b) {
                child->level = q.front()->level + 1;
                q.push(child);
            }
        }
        top.push_back(q.front());
        q.pop();
    }
}

void evaluate(vector<FormulaNode *> &top) {
	int i;

    for(FormulaNode *next_f : top) {
        for(FormulaNode *parent : next_f->parents) {
            if(parent->result == -1 || parent->level >= next_f->level) {
			    next_f->result = -1;
				break;
			}
		    next_f->result += (parent->result == next_f->value) ? 1 : 0;
        }
    }
}
