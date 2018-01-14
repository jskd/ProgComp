import sys

def read_file(filename):
    res = []
    with open(filename) as f:
        for line in f:
            res.append(line.strip())
    return res

words = []
dic = read_file(sys.argv[1])

for a in range(2, len(sys.argv)):
    words.append(sys.argv[a])


for k in dic:
    for g in words:
        if ''.join(sorted(k)) == ''.join(sorted(g)):
            print(k)
