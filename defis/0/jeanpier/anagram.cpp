
#include <iostream>
#include <fstream>
#include <algorithm>
#include <memory>
#include <string>
#include <vector>
#include <map>

static const int MIN_ARGC   = 3;
static const int START_ARGV = 2;

/*
*   I need this structure to store the words I want to use to get their anagrams
*/
struct WordParams
{
    const int NB_WORDS;
    std::unique_ptr<std::string[]> wparam;
    std::unique_ptr<std::string[]> sorted_wparam;
    std::map<std::string, std::vector<std::string>> anagrams;

    WordParams(char ** str, int sz);
};

WordParams::WordParams(char ** str, int sz)
    : NB_WORDS(sz), wparam(new std::string[sz]), sorted_wparam(new std::string[sz]),
    anagrams()
{
    int j = 0;
    std::string _stmp;
    for(int i = 2; i < sz; i++)
    {
        _stmp = str[i];
        wparam[j] = _stmp;
        anagrams[_stmp] = std::vector<std::string>();

        std::sort(_stmp.begin(), _stmp.end());
        sorted_wparam[j] = _stmp;
        ++j;
    }
}

// Main macros and functions
#define ERROR(argv) std::cerr << "usage: " << argv[0] << " dict_file [words]\n"

bool areParamOK(int argc, char ** argv) noexcept;
int  anagrams(int argc, char ** argv);

void generateAnagrams(WordParams& wp, const std::string& file);
void printAnagrams(const WordParams& wp) noexcept;


int main(int argc, char ** argv)
{
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
    WordParams wp(argv, argc);
    generateAnagrams(wp, argv[1]);
    printAnagrams(wp);
    return 0;
}


void generateAnagrams(WordParams& wp, const std::string& file)
{
    std::ifstream _infile(file);

    if(_infile.is_open())
    {
        std::string _line;

        while(std::getline(_infile, _line))
        {
            std::string _sline(_line);

            // lowercase
            std::transform(_sline.begin(), _sline.end(), _sline.begin(),
                           [](unsigned char c){ return std::tolower(c); } );

            std::sort(_sline.begin(), _sline.end());

            // test if the word is an anagram of on of the keys
            for(int i = 0 ; i < wp.NB_WORDS; i++)
            {
                auto& _key = wp.wparam[i];
                auto& _sorted_key = wp.sorted_wparam[i];

                if(_sorted_key == _sline)
                {
                    wp.anagrams[_key].push_back(_line);
                    break;
                }
            }
        }
    }
}

void printAnagrams(const WordParams& wp) noexcept
{
    for(const auto& item: wp.anagrams)
    {
        std::cout << item.first << ":\n";
        for(const std::string& s: item.second)
        {
            std::cout << s << "\n";
        }
        std::cout << "\n";
    }
}
