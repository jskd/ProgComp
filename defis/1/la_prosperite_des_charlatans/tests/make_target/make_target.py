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

        # BEGIN OF SCRIPT
        test_result = True

        try:
            bin_dir = tested_path + "/" + "bin/"

            if not os.path.isdir(bin_dir):
                raise Exception("Error: bin/ does not exist.")

            with open(tested_path + "/" + "Makefile", "r"): pass
            cmd = ["make", "--directory=" + tested_path]
            out = check_output(cmd, stderr=STDOUT, timeout=30).decode("utf-8")
            matchObj = re.match(r"make\[1\]*", out)

            if matchObj:
                test_result = False

        except:
            out = "False"
            test_result = False

        print(test_result) # True for PASS, False for FAIL
        # END OF SCRIPT

    else:
        print("Usage: files_presence.py <expected_file>")
