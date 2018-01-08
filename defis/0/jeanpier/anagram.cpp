
#include <iostream>
#include <fstream>
#include <algorithm>
#include <string>
#include <vector>
#include <map>

#define MIN_ARGC 3

void readDict(const std::string& file, std::vector<std::string>& v);

int main(int argc, char ** argv)
{
    if(argc < MIN_ARGC)
    {
        std::cerr << "usage: " << argv[0] << "dict_file [words]\n";
        return -1;
    }

    std::vector<std::string> file_content;
    std::map<std::string, std::vector<std::string> > anagrams;

    readDict(std::string(argv[1]), file_content);

    for(std::string& w: file_content)
    {
        std::string key = w;
        std::sort(key.begin(), key.end());
        std::vector<std::string> & v = anagrams[key];
        v.push_back(w);
    }

    /*for(int i = 2; i < argc; i++)
    {   // TODO
        std::string word_to_search = argv[i];

        if(anagrams.get())
    }*/

    return 0;
}


void readDict(const std::string& file, std::vector<std::string>& v)
{
    std::ifstream infile(file);
    std::string line;

    while(std::getline(infile, line))
    {
        v.push_back(line);
    }
}
