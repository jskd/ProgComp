#include <iostream>
#include <functional>

#include "parser.hpp"
#include "formulanode.hpp"
#include "rtree.hpp"

using namespace std;

int unsafe_main(int argc, char *argv[]) {
    if(argc != 5) {
        cerr << "Are you kidding me? Something wrong with your arguments.\n";
        return 1;
    }
    Parser p(argv[1]);
    vector<FormulaNode *> rs, top;
    INode *head;

#ifndef CHICKEN_MODE
    init_opencl();
#endif
    head = roots(p, rs);
    normalize(rs, top);
    head->preval(p);
    // Apparently not so superior:
    // superior_preval(p);
    evaluate(top);
    write_view(p, argv[3]);

#ifdef VERBOSE
    cout << "All formulas:\n";
    head->foreach([](FormulaNode &f) { cout << f; });

    cout << "\nRoots:\n";
    for(FormulaNode *f : rs) {
        cout << *f;
    }
#endif

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
