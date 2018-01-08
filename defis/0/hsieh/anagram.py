import sys
from collections import deque

arg_length = len(sys.argv)
if(arg_length>2):
    dic = {}
    file_path = sys.argv[1]
    print("Reading file: "+file_path)
    f = open(file_path, 'r')
    line_count = 0
    for line in f:
        one_line = line.strip()
        dic[one_line] = line_count
        #print("["+one_line+"]="+str(line_count))
        line_count+=1
    f.close()
    for i in range (2, arg_length):
        input_txt = sys.argv[i]
        #print("Checking "+input_txt)
        if(input_txt in dic):
            print(input_txt)
        chars = list(input_txt)
        items = deque(chars)
        for itr in range(1, len(chars)):
            items.rotate(1)
            rotated_txt = "".join(list(items))
            #print("Rotated txt="+rotated_txt)
            if(rotated_txt in dic):
                print(roated_txt)
        

else:
    print("Not enough parameters")
