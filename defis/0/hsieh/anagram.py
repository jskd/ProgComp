# Usage:
#    python3 anagram.py <dictionary_file> <vocabulary1> <vocabulary2>...
# Example:
#    python3 anagram.py ../american-english foo bar etc

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
        dic[one_line] = line_count # Put key -> value into dictionary as vocabulary -> line_number
        #print("["+one_line+"]="+str(line_count))
        line_count+=1
    f.close()
    for i in range (2, arg_length):
        input_txt = sys.argv[i]
        #print("Checking "+input_txt)
        if(input_txt in dic): # If current word is in the dictionary, print out as result
            print(input_txt)
        chars = list(input_txt) # Split word into char array
        items = deque(chars) # Use deque in python library to recombine the chars
        for itr in range(1, len(chars)):
            items.rotate(1) # Rotate deque by one char to the right. So that "etc" becomes "cet"
            rotated_txt = "".join(list(items)) # Combine all the rotated chars to string
            #print("Rotated txt="+rotated_txt)
            if(rotated_txt in dic):
                print(roated_txt)
        

else:
    print("Not enough parameters")
