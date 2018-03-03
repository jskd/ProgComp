#include <iostream>
#include <functional>

#include "parser.hpp"
#include "cell.hpp"
#include "rtree.hpp"

using namespace std;

int unsafe_main(int argc, char *argv[]) {
    if(argc < 2) {
        cerr << "File, please.\n";
        return 1;
    }
    Parser p(argv[1]);
    vector<FormulaNode *> rs, top;
    INode *head;

    init_opencl();
    head = roots(p, rs);
    normalize(rs, top);
    head->preval(p);
    //superior_preval(p);
    evaluate(top);

    cout << "All formulas:\n";
    head->foreach([](FormulaNode &f) { cout << f; });

    cout << "\nRoots:\n";
    for(FormulaNode *f : rs) {
        cout << *f;
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
