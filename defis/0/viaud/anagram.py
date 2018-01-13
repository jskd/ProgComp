# Viaud
# Python 2

import sys

def open_file(filename):
    f = open(filename, 'r')
    return f

def populate_dict(word, dictionary):
    for char in word:
        # On pourrait utiliser une ternaire ici mais je laisse volontairement
        # comme ca pour la lisibilite !
        if char in dictionary:
            dictionary[char] += 1
        else:
            dictionary[char] = 1

def are_anagrams(word1, word2):
    wd1 = {}
    wd2 = {}
    populate_dict(word1, wd1)
    populate_dict(word2, wd2)
    return cmp(wd1, wd2) == 0

# Plus lent hehe !
def are_anagrams_2(word1, word2):
    return sorted(list(word1.lower())) == sorted(list(word2.lower()))

def compute(file, anagrams):
    for line in file:
        sanitized = line.replace('\n', '')
        for word in sys.argv[2:]:
            if are_anagrams(word.lower(), sanitized.lower()):
                if word in anagrams:
                    anagrams[word].append(sanitized)
                else:
                    anagrams[word] = [sanitized]

# Affiche les anagrams a partir du dictionnaire
# Cles = Mots, Valeurs = liste de mots
def print_anagrams(anagrams):
    if len(anagrams) == 0:
        print "No anagrams found!"
    else:
        for word in anagrams:
            print word + ":"
            for anagram in anagrams[word]:
                print anagram

try:
    filename = sys.argv[1]
    f = open_file(filename)
    anagrams = {}
    compute(f, anagrams)
    print_anagrams(anagrams)
except:
    sys.exit("Please provide a valid dictionary.")
