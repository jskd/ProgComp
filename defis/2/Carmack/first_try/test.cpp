#include <iostream>

#include "parser.hpp"

int main(int argc, char *argv[]) {
    cell c;
    Parser p("data.csv");

    cout << p.next_cell(&c) << " " << (c.type == None) <<
        " " << c.value << endl;
    cout << p.next_cell(&c) << " " << (c.type == None) <<
        " " << c.value << endl;
    cout << p.next_cell(&c) << " " << (c.type == None) <<
        " " << c.value << endl;
    cout << p.next_cell(&c) << " " << (c.type == None) <<
        " " << c.v << endl;
    cout << p.next_cell(&c) << " " << (c.type == None) <<
        " " << c.v << endl;
    cout << p.next_cell(&c) << " " << (c.type == None) <<
        " " << c.value << endl;
    return 0;
}
