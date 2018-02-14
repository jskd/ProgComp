#include <iostream>
#include <vector>
using namespace std;

int main(int argc, char *argv[]) {
    
	vector<unsigned> a;
	cout << a.size() << endl;
	a.push_back(14);
	a.push_back(34);
	a.push_back(56);
	cout << a.size() << endl;
	return 0;
}
