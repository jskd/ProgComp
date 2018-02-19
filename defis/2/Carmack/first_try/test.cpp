#include <iostream>
#include <functional>

#include "parser.hpp"
#include "cell.hpp"
#include "rtree.hpp"

int main(int argc, char *argv[]) {
    if(argc < 2) {
        cerr << "File, please.\n";
        return 1;
    }
    Parser p(argv[1]);
    vector<formula *> rs, top;
    INode *head = roots(p, rs);

    //normalize(rs, top);
    //p.reposition();
    //head->preval(p);
	//evaluate(top);

    cout << "All formulas:\n";
    head->foreach([](formula &f) { cout << f; });

    cout << "\nRoots:\n";
    vector<formula *>::iterator i;

    for(i = rs.begin(); i != rs.end(); ++i) {
        cout << **i;
    }
}
