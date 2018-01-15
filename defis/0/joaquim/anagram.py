import sys

# Retourne la liste des mots du dictionnaire
def getWordList(dico_path):
	words = []
	with open(dico_path) as f:
		for line in f.readlines():
			stripped = line.strip("\n").lower()
			words.append(stripped)
	return words

# Test si word1 et word2 sont des anagram entre eux
def isAnagramOf(word1, word2):
	w1 = "".join(sorted(list(word1)))
	w2 = "".join(sorted(list(word2)))

	if w1 == w2: return True
	return False

if __name__ == "__main__":
	if len(sys.argv) > 1:
		dico_path = sys.argv[1]
		words_args = sys.argv[2::]

		# Création du dictionnaire des anagrammes
		dico_anagram = dict((el.lower(),[]) for el in words_args)
		# Récupération de la liste des mots du dictionnaire
		dico_words = getWordList(dico_path)

		for dico_entry in dico_words:
			for word in dico_anagram.keys():
				if isAnagramOf(word, dico_entry):
					dico_anagram[word] += [dico_entry]

		for k, v in dico_anagram.items():
			print(k+":")
			for value in v:
				print(value)
	else:
		print("Usage: anagram.py <dico> <words>")
