#include <iostream>

using namespace std;

int main(int argc, char** argv){
  if(argc >= 4){
    char* data_path = argv[0];
    char* user_path = argv[1];
    char* view0_path = argv[2];
    char* changes_path = argv[3];

    cout << "Compute many things..." << endl;
  }
  else{
    cout << "Usage: ws <data> <user> <view0> <changes>" << endl;
  }
  exit(0);
}
