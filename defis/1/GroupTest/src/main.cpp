#include <iostream>
#include <fstream>
#include <string>

using namespace std;

int main(int argc, char *argv[]){
  if(argc >= 4){
    string data_path = argv[1];
    string user_path = argv[2];
    string view0_path = argv[3];
    string changes_path = argv[4];

    ofstream out_view0(view0_path);
    out_view0 << "1,1,1,3" << endl;
    out_view0.close();

    ofstream out_changes(changes_path);
    out_changes << "" << endl;
    out_changes.close();

  }
  else{
    cout << "Usage: ws <data> <user> <view0> <changes>" << endl;
  }
  exit(0);
}
