import sys

if __name__ == "__main__":
	if len(sys.argv) > 1:
		words = sys.argv[1::]

		print("Mots = " + str(words))

		for word in words:
			print(word + ":")
			words_letters = list(set(word))

			with open("american-english") as f:
				for line in f:
					stripped = line.strip("\n").lower()
					current = list(stripped)
					current_letters = list(set(current))

					if (len(word) == len(current)) and (words_letters == current_letters):
						print(current)

				f.close()



	else:
		print("Usage: defis0.py <words>")