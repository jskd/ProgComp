import sys

# Ché pas se ke je kode!

def isAnagam(str1, str2):
    for s in list(str1):
        if s not in list(str2):
            return False
    return True

if(len(sys.argv) < 2):
    print("Wesh, usage: anagram mydict foo bar")
    sys.exit()

with open(sys.argv[1]) as f:
    for word in f.readlines():
        for anagram in sys.argv[2::]:
            if isAnagam(word, anagram):
                print(word)
