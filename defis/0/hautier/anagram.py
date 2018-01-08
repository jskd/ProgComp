import sys

def read_file(filename):
    res = []
    with open(filename) as f:
        for line in f:
            res.append(line.strip())
    return res

def permut(word,i, j):
    liststring = []
    for a in word:
        liststring+=a
    tmp = liststring[j]
    liststring[j] = liststring[i]
    liststring[i] = tmp

    return ''.join(liststring)

def genanagram(w):
    anagramlist = set() 
    for a in range(len(w)):
        for b in range(len(w)):
            anagramlist.add(permut(w, a, b))
    return anagramlist

words = []
dic = read_file(sys.argv[1])

for a in range(2, len(sys.argv)):
    words.append(list(genanagram(sys.argv[a])))

for k in dic:
    for g in words:
        if k in g:
            print(k)