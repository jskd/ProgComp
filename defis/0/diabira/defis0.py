# coding: utf-8

import sys

def anagram(dictionnaire,mots):
    dico = open(dictionnaire,"r").readlines()
    tmp = []
    for mot in mots:
        print(mot+":")
        for d in dico:
            if sorted(mot)== sorted(d.replace('\n','')):
                tmp.append(d.replace('\n',''))
        affiche(sorted(tmp))
    
def affiche(l):
    for e in l:
        print(e)

if __name__ == "__main__":
    if len(sys.argv) < 3:
        print("need at least two args")
    else:
        anagram(sys.argv[1],sys.argv[2:])