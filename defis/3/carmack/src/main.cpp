#include <iostream>

#include "parser.hpp"

using namespace std;

extern "C" {
    #include "formula.h"
}

void fst_pass(Parser &p) {
    cell c;

    p.reposition();

    while(!p.eof) {
        p.next_cell(&c);

        if(c.type == Formula)
            new_formula(c.x1, c.y1, c.x2, c.y2, c.v);
    }
}

void preval(Parser &p) {
    cell c;
    int x = 0, y = 0;

    p.reposition();

    while(!p.eof) {
        p.next_cell(&c);

        if(c.type == Value)
        cout << "x = " << x << " y = " << y << " value = " << c.value << endl;
        if(c.type == Value && push_value(x, y, c.value))
            launch_calculation();

        if(p.eol) {
            x++;
            y = 0;
        }
        else
            y++;
    }

    launch_calculation();
}

int unsafe_main(int argc, char *argv[]) {
    if(argc != 5) {
        cerr << "Are you kidding me? Something wrong with your arguments.\n";
        return 1;
    }
    Parser p(argv[1]);
    int id = 0;
    formula *f;

    init_opencl();
    fst_pass(p);
    preval(p);


    while((f = get_formula(id))) {
        cout << "# Formula " << id << " #\n" <<
            "\tx1 = " << f->x1 << ", y1 = " << f->y1 <<
            ", x2 = " << f->x2 << ", y2 = " << f->y2 <<
            ", v = " << f->v << "\n\tresult: " << f->res << endl;
        id++;
    }

    return 0;
}

int main(int argc, char *argv[]) {
    try {
        return unsafe_main(argc, argv);
    }
    catch(string e) {
        cout << e << endl;
    }
    return 0;
}
