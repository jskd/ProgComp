#! /usr/bin/python3

#create a file de big_mama that is ? mb.

import sys, csv, random, os
from os.path import getsize

max_rows = 1000

if len(sys.argv) != 2 or not sys.argv[1].isdigit():
	print("The arg1 must be a int")
	exit(0)

file_name = "bigmama_{}.csv".format(sys.argv[1])
total_lines = int(int(sys.argv[1])*1024*1024/3590)
#return a int between a and b
def between(a,b):
	return random.randint(a,b)

#return a formula
def get_formula():
	r1 = between(0,max_rows-7)
	c1 = between(0,total_lines-7)
	r2 = between(r1+1,max_rows)
	c2 = between(c1+1,total_lines)
	#return "\"=#("+str(r1)+","+str(c1)+","+str(r2)+","+str(c2)+","+str(between(0,255))+")\""
	return "=#("+str(r1)+","+str(c1)+","+str(r2)+","+str(c2)+","+str(between(0,255))+")"

#return a table with intergers and a formula
def get_row():
	result = []
	for i in range(0,max_rows):
		result.append(str(between(0,255)))
	result[between(0,max_rows-1)] = get_formula()
	return result

with open(file_name,"w+",newline='') as csv_file:
	writer = csv.writer(csv_file,delimiter=";")

	for i in range(0,total_lines):
		writer.writerow(get_row())
