import sys

# Must contain at least 3 arguments
if len(sys.argv) < 3 :
	print("Not enough arguments")
	sys.exit()

# Dictionairy file
dictionairy_file = open(sys.argv[1], "r")


if dictionairy_file : 
	print("Not a dictionairy file")
	sys.exit()


#Test if a string is annagramm of other
def is_annagramm(str1, str2):
	if not str1 and not str2:
		return True
	elif str1 and str2:
		if str1[0] in str2:
			is_annagramm(str1[:1], str2.strip(str1[0]))
		else:
			return False
	else:
		return False


#Return a list of all anagramms in a file
def listOfAnnagram(file,str):
	lst = []
	for word in file:
		if is_annagramm(str,word):
			lst.add(word)
	return lst



#Read all words given as arguments
for word in sys.argv[:2]:
	lst = listOfAnnagram(dictionairy_file,word)
	if lst :
		print(word + ":\n")
		lst.sort()
		print(lst)
