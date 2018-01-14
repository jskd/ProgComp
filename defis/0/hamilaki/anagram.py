#!/usr/bin/env python3

import sys
import os.path


# Must contain at least 3 arguments
if len(sys.argv) < 3 :
	print("Not enough arguments")
	print("use : ", sys.argv[0], "<dictionary>", " <list of words>")
	sys.exit()

# Argument 1 must be a file
if not os.path.isfile(sys.argv[1]): 
	print("Error no dictionary file")
	sys.exit()

# Dictionary
dictionairy = sys.argv[1]

#Test if a string is anagram of other
def is_anagram(str1, str2):
	return sorted(str1.lower()) == sorted(str2.lower())

#Return a list of all anagrams of a string in a file
def list_anagrams(filename,str):
	lst = []
	stripped_str = str.rstrip()
	with open(filename, "r") as file:
		for word in file:
			stripped_word = word.rstrip()
			if is_anagram(stripped_str,stripped_word):
				lst.append(stripped_word)
	return lst

# For each argument print list of anagrams
for word in sys.argv[2:]:
	lst = list_anagrams(dictionairy,word)
	lst.sort()
	print(word + ":")
	print("\n".join(lst))
