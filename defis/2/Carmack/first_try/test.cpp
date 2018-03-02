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
    vector<formula *> rs, top;
    INode *head = roots(p, rs);

    normalize(rs, top);
    head->preval(p);
    evaluate(top);

    cout << "All formulas:\n";
    head->foreach([](formula &f) { cout << f; });

    cout << "\nRoots:\n";
    for(formula *f : rs) {
        cout << *f;
    }
}

int main(int argc, char *argv[]) {
    try {
        unsafe_main(argc, argv);
    }
    catch(string e) {
        cout << e << endl;
    }
    return 0;
}
