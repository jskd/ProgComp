import sys

def read_file(filename):
    res = []
    with open(filename) as f:
        for line in f:
            res.append(line.strip())
    return res

if(len(sys.argv) < 3):
    print("Erreur")
    sys.exit(-1)

words = sys.argv[2::] 
dic = read_file(sys.argv[1])

for k in dic:
    for g in words:
        if ''.join(sorted(k)) == ''.join(sorted(g)):
            print(k)
