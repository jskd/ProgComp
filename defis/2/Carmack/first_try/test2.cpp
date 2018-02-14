#include <iostream> 
#include "structures.hpp"
#include "parser.hpp"
#include <sstream>

void test_parse(){
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
    return;
}
void print_area(area a){
	printf("(%d, %d, %d, %d)\n", a.x1, a.y1, a.x2, a.y2);
}

int main(int argc, char *argv[]) {
    test_parse();
	cell c;
	c.type = Value;
	c.value = 5;
	node b;
	b.f = c;
	c.value = 6;
	cout << b.f.value << endl;
	graph g;
	g.push_back(&b);
	vector<int> v;
	cout << (v.end() - v.begin()) << endl;
	area a, a2;
	fill_a(a, 3, 4, 5, 6);
	fill_a(a2, 3, 5, 5, 6);
	print_area(a2);
	cout << equal_a(a, a2) << endl;
	a2.y1 = 4;
	print_area(a2);
	cout << equal_a(a, a2) << endl;
	cout << is_in_a(a, 3, 4) << endl;
	cout << is_in_a(a, 4, 4) << endl;
	cout << is_in_a(a, 4, 5) << endl;
	cout << is_in_a(a, 5, 6) << endl;
	cout << is_in_a(a, 2, 4) << endl;
	cout << is_in_a(a, 6, 4) << endl;
	cout << is_in_a(a, 4, 3) << endl;
	cout << is_in_a(a, 4, 7) << endl;
	fill_a(a, 1, 4, 7, 8);
	fill_a(a2, 5, 3, 4, 16);
	print_area(least_upper_a(a, a2));
	cerr << "Error!\n" << endl;
	return 0;
	
}
