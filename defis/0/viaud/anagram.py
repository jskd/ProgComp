# Code moche
# Typiquement (si on avait plus de temps...)
# Je testerai directement tous les mots (argv) pour une line du fichier
# Plutot que d'iterer

import sys
from collections import Counter

def open_file(filename):
    f = open(filename, 'r')
    return f

def is_anagram(word1, word2):
    return Counter(word1) == Counter(word2)

def is_anagram2(word1, word2):
    return sorted(list(test.lower())) == sorted(list(original.lower())) 

def is_anagram3(word1, word2):
    for l1 in range(0, len(word1), 1):
        if l2 in range (0, len(word2), 1):
            return True
        else:
            return False

def anagrams(file, word):
    for line in file:
        cleaned_line = line.replace('\n', '')
        if is_anagram3(word,cleaned_line):
            print cleaned_line 

def print_anagrams(file):
    for word in sys.argv[2:]:
        print word + ":" 
        anagrams(file, word)
        file.seek(0)

try:
    filename = sys.argv[1]
    f = open_file(filename)
    print_anagrams(f)
except:
    sys.exit("Please provide a valid dict.")    
