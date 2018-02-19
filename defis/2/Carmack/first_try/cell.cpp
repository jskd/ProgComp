#include <queue>

#include "cell.hpp"

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
    vector<formula *>::iterator i, j;

    for(i = roots.begin(); i != roots.end(); ++i) {
        for(j = (*i)->children.begin(); j != (*i)->children.end(); ++j) {
            (*j)->level = 1;
            q.push(*j);
        }
        top.push_back(*i);
    }

    while(!q.empty()) {
        for(j = q.front()->children.begin();
            j != q.front()->children.end(); ++j) {
            if((*j)->level != -1)
                continue;
            (*j)->level = q.front()->level + 1;
            q.push(*j);
        }
        top.push_back(q.front());
        q.pop();
    }
}
