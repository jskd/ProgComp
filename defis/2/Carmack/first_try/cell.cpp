#include <queue>

#include "cell.hpp"

using namespace std;

formula::formula() { }

formula::formula(int x1, int y1, int x2, int y2, int v) {
    bb = area(x1, y1, x2, y2);
    value = v;
}

formula::formula(cell &c) : formula(c.x1, c.y1, c.x2, c.y2, c.v) { }

ostream &operator<<(ostream &out, const formula &f) {
    out << "# Formula #\n" << "\tposition: x = " << f.p.x << ", y = " <<
        f.p.y << "\n\tx1 = " << f.bb.x1 << ", y1 = " << f.bb.y1 <<
        ", x2 = " << f.bb.x2 << ", y2 = " << f.bb.y2 << ", v = " <<
        f.value << "\n\tactual result: " << f.result << "\n\tlevel = " <<
        f.level << endl;
}

void normalize(vector<formula *> &roots, vector<formula *> &top) {
    queue<formula *> q;

    for(formula *root : roots) {
        for(formula *child : root->children) {
            child->level = 1;
            q.push(child);
        }
        // Zero-level formulae are already calculated,
		// so there is no reason to add them to top.
        // top.push_back(child);
    }

    while(!q.empty()) {
        for(formula *child : q.front()->children) {
            child->level = q.front()->level + 1;
            q.push(child);
        }
        top.push_back(q.front());
        q.pop();
    }
}

void evaluate(vector<formula *> &top) {
	int i;

    for(formula *next_f : top) {
        for(formula *parent : next_f->parents) {
            if(parent->result == -1 || parent->level >= next_f->level) {
			    next_f->result = -1;
				break;
			}
		    next_f->result += (parent->result == next_f->value) ? 1 : 0;
        }
    }
}
