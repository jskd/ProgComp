#include <queue>

#include "formulanode.hpp"

using namespace std;

FormulaNode::FormulaNode(int x1, int y1, int x2, int y2, int v) :
    formula_id(new_formula(x1, y1, x2, y2, v)) { }

FormulaNode::FormulaNode(cell &c) :
    FormulaNode(c.x1, c.y1, c.x2, c.y2, c.v) { }

Area FormulaNode::bb() {
    formula *f = get_formula(formula_id);

    return Area(X1(f), Y1(f), X2(f), Y2(f));
}

int FormulaNode::value() {
    return V(get_formula(formula_id));
}

void FormulaNode::value(int v) {
    V(get_formula(formula_id)) = v;
}

int FormulaNode::result() {
    return R(get_formula(formula_id));
}

void FormulaNode::result(int r) {
    R(get_formula(formula_id)) = r;
}

int FormulaNode::level() {
    return LEVEL(get_formula(formula_id));
}

void FormulaNode::level(int l) {
    LEVEL(get_formula(formula_id)) = l;
}

ostream &operator<<(ostream &out, FormulaNode &f) {
    Area f_bb = f.bb();

    out << "# Formula #\n" << "\tposition: x = " << f.p.x << ", y = " <<
        f.p.y << "\n\tx1 = " << f_bb.x1 << ", y1 = " << f_bb.y1 <<
        ", x2 = " << f_bb.x2 << ", y2 = " << f_bb.y2 << ", v = " <<
        f.value() << "\n\tactual result: " << f.result() << "\n\tlevel = " <<
        f.level() << endl;
    return out;
}

void normalize(vector<FormulaNode *> &roots, vector<FormulaNode *> &top) {
    queue<FormulaNode *> q;
    bool b;

    for(FormulaNode *root : roots) {
        for(FormulaNode *child : root->children) {
            child->level(1);
            q.push(child);
        }
        // Zero-level formulas are already calculated,
        // so there is no reason to add them to top.
    }

    while(!q.empty()) {
        for(FormulaNode *child : q.front()->children) {
            if(child->level() != -1)
                continue;
            b = true;
            for(FormulaNode *parent : child->parents) {
                if(parent->level() == -1) {
                    b = false;
                    break;
                }
            }
            if(b) {
                child->level(q.front()->level() + 1);
                q.push(child);
            }
        }
        top.push_back(q.front());
        q.pop();
    }
}

void evaluate(vector<FormulaNode *> &top) {
    for(FormulaNode *next_f : top) {
        for(FormulaNode *parent : next_f->parents) {
            if(parent->result() == -1 || parent->level() >= next_f->level()) {
			    next_f->result(-1);
				break;
			}

		    next_f->result(next_f->result() +
                           ((parent->result() == next_f->value()) ? 1 : 0));
        }
    }
}

void superior_preval(Parser &p) {
    point pos;
    cell c;

    pos.x = pos.y = 0;
    p.reposition();

    while(!p.eof) {
        p.next_cell(&c);
        switch(c.type) {
        case None: case Formula: break;
        case Value:
            if(push_value(pos.x, pos.y, c.value))
                launch_calculation();
        }

        if(p.eol) {
            pos.y = 0;
            pos.x++;
        }
        else
            pos.y++;
    }

    launch_calculation();
}

int write_view(Parser &p, string view) {
    cell c;
    point pos;
    int f_id = 0;
    ofstream fs(view);

    if(!fs)
        return -1;

    pos.x = pos.y = 0;
    p.reposition();

    while(!p.eof) {
        p.next_cell(&c);
        switch(c.type) {
        case None:
            break;
        case Value:
            fs << c.value << ";";
            break;
        case Formula:
            if(R(get_formula(f_id)) == -1)
                fs << "R;";
            else
                fs << R(get_formula(f_id)) << ";";
            f_id++;
            break;
        }

        if(p.eol && !p.eof) {
            fs << endl;
            pos.y = 0;
            pos.x++;
        }
        else
            pos.y++;
    }

    fs.close();
    return 0;
}
