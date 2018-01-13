
#include <iostream>
#include <fstream>
#include <algorithm>
#include <memory>
#include <string>
#include <vector>
#include <map>

/*
*   I need this structure to store the words I want to use to get their anagrams
*/
struct WordParams
{
    std::unique_ptr<std::string[]> wparam;
    std::unique_ptr<std::string[]> sorted_wparam;
};

void wp_init(char ** str, int sz);

// Main macros and functions
#define ERROR(argv) std::cerr << "usage: " << argv[0] << " dict_file [words]\n"

void readDict(const std::string& file, std::vector<std::string>& v);

bool areParamOK(int argc, char ** argv) noexcept;
int  anagrams(int argc, char ** argv);

void generateAnagrams();
void printAnagrams() noexcept;

static const int MIN_ARGC   = 3;
static const int START_ARGV = 2;


int main(int argc, char ** argv)
{

    /*std::vector<std::string> file_content;
    std::map<std::string, std::vector<std::string> > anagrams;

    readDict(std::string(argv[1]), file_content);

    for(std::string& w: file_content)
    {
        std::string key = w;
        std::sort(key.begin(), key.end());
        std::vector<std::string> & v = anagrams[key];
        v.push_back(w);
    }*/

    /*for(int i = 2; i < argc; i++)
    {   // TODO
        std::string word_to_search = argv[i];

        if(anagrams.get())
    }*/

    return areParamOK(argc, argv) ? anagrams(argc, argv) : -1;
}


bool areParamOK(int argc, char ** argv) noexcept
{
    if(argv == NULL)
    {
        std::cerr << "error: null argv\n";
        ERROR(argv);
        return false;
    }

    if(argc < MIN_ARGC)
    {
        ERROR(argv);
        return false;
    }

    for(int i = 1; i < argc; i++)
    {
        if(argv[i] == NULL)
        {
            std::cerr << "error: null argument at " << i << "\n";
            ERROR(argv);
            return false;
        }
    }

    return true;
}

/**
*   @note I don't need to check argc and argv (already checked). See areParamOK()
*/
int anagrams(int argc, char ** argv)
{
    WordParams wp {
        std::unique_ptr<std::string[]>(new std::string[argc - START_ARGV]),
        std::unique_ptr<std::string[]>(new std::string[argc - START_ARGV])
    };

    int j = 0;
    for(int i = START_ARGV; i < argc; i++)
    {
        std::string _stmp(argv[i]);
        wp.wparam[j] = _stmp;

        std::sort(_stmp.begin(), _stmp.end());
        wp.sorted_wparam[j] = _stmp;
        std::cout << "tmp: " << wp.wparam[j] << "\n";
        std::cout << "sorted tmp: " << wp.sorted_wparam[j] << "\n";
        ++j;
    }

    generateAnagrams();
    printAnagrams();
    return 0;
}


void generateAnagrams()
{

}


void printAnagrams() noexcept
{

}

// WordParams implementation
void wp_init(char ** str, int sz)
{
    
}

// Dump

void readDict(const std::string& file, std::vector<std::string>& v)
{
    std::ifstream infile(file);
    std::string line;

    while(std::getline(infile, line))
    {
        v.push_back(line);
    }
}
