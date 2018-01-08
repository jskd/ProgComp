import sys

def isAnagram(str1, str2):
    for s in list(str1.lower()):
        if s not in list(str2.lower()):
            return False
    return True

if(len(sys.argv) < 2):
    print("Usage: anagram mydict foo bar")
    sys.exit()

with open(sys.argv[1]) as f:
    for word in f.readlines():
        for anagram in sys.argv[2::]:
            if isAnagram(word, anagram):
                print(word)
