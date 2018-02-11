#include <iostream>

#include "parser.hpp"
#include <sstream>

int main(int argc, char *argv[]) {
    cell c;
    Parser p("data.csv");

	stringstream s;
	bool b;
	while(true){
		b = p.next_cell(&c);
		switch(c.type){
			case(None):
			s << "None; ";
			break;
			case(Value):
			s << "Value: " << c.value << "; ";
			break;
			case(Formula):
			s << "Formula: " << c.x1 << " " << c.y1 << " " << 
			c.x2 << " " << c.y2 << " " << c.v << "; ";
			break;
		}
		cout << s.str() << endl;
		s.str(std::string());
		if (b)
			cout << "End of line!" << endl;
		if (p.eof)
			break;
	}
	cout << "End of file!" << endl;
    return 0;
}
