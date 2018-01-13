
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
    const int NB_WORDS;
    std::unique_ptr<std::string[]> wparam;
    std::unique_ptr<std::string[]> sorted_wparam;
    std::map<std::string, std::vector<std::string>> anagrams;

    //WordParams(char ** str, int sz);
};
/*
WordParams::WordParams(char ** str, int sz)
{

}*/

void wp_init(WordParams& wp, char ** str, int sz);

// Main macros and functions
#define ERROR(argv) std::cerr << "usage: " << argv[0] << " dict_file [words]\n"

void readDict(const std::string& file, std::vector<std::string>& v);

bool areParamOK(int argc, char ** argv) noexcept;
int  anagrams(int argc, char ** argv);

void generateAnagrams(WordParams& wp, const std::string& file);
void printAnagrams(const WordParams& wp) noexcept;

static const int MIN_ARGC   = 3;
static const int START_ARGV = 2;


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
    const int SZ = argc - START_ARGV;
    WordParams wp {
        SZ,
        std::unique_ptr<std::string[]>(new std::string[SZ]),
        std::unique_ptr<std::string[]>(new std::string[SZ]),
        std::map<std::string, std::vector<std::string>>()
    };

    wp_init(wp, argv, argc);
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

            for(int i = 0 ; i < wp.NB_WORDS; i++)
            {
                auto& _sorted_key = wp.sorted_wparam[i];

                if(_sorted_key == _sline)
                {
                    wp.anagrams[wp.wparam[i]].push_back(_line);
                    break;
                }
            }
        }
    }
}


void printAnagrams(const WordParams& wp) noexcept
{
    for(auto& item: wp.anagrams)
    {
        std::cout << item.first << ":\n";
        for(const std::string& s: item.second)
        {
            std::cout << s << "\n";
        }
    }
}

// WordParams implementation
void wp_init(WordParams& wp, char ** str, int sz)
{
    int j = 0;
    for(int i = START_ARGV; i < sz; i++)
    {
        std::string _stmp(str[i]);
        wp.wparam[j] = _stmp;
        wp.anagrams[_stmp] = std::vector<std::string>();

        std::sort(_stmp.begin(), _stmp.end());
        wp.sorted_wparam[j] = _stmp;
        ++j;
    }
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
