#!/usr/bin/env python3

import json
import os,sys
from subprocess import *
from os import listdir
import os.path
from os.path import isfile, join
import re


if __name__ == "__main__":
    if len(sys.argv) >= 1:
        path = os.path.dirname(sys.argv[0]) + "/"
        tested_path = sys.argv[1]

        # INIT DATA
        with open(path  + "infos.json", 'r') as data:
            test_data = json.load(data)
            test_info = test_data["infos"][0]

        test_result = True

        try:
            with open(tested_path + "/" + "view0.csv", "r") as user_file:

                for line in user_file.readlines():
                    line_splited = line.strip("\n").split(" ")

                    matchObj = re.match(r"(\d.*|P)(,\d.*|P)*", line)

                    if not matchObj:
                        test_result = False
                        break


        except: test_result = False
        # BEGIN OF SCRIPT
        print(test_result) # True for PASS, False for FAIL
        # END OF SCRIPT

    else:
        print("Usage: files_presence.py <expected_file>")
