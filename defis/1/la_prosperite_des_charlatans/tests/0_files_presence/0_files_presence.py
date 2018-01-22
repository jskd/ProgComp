#!/usr/bin/env python3

import json
import os,sys
from subprocess import *
from os import listdir
import os.path
from os.path import isfile, join


if __name__ == "__main__":
    if len(sys.argv) >= 1:
        path = os.path.dirname(sys.argv[0]) + "/"
        tested_path = sys.argv[1]

        # INIT DATA
        with open(path  + "infos.json", 'r') as data:
            test_data = json.load(data)
            test_info = test_data["infos"][0]

        # BEGIN OF SCRIPT
        expected_path = join(path, "expected/" + test_info["expected"])

        with open(expected_path, 'r') as expected_files:
            expected_liste = sorted(expected_files.read().strip("\n").split("\n"))

        files_liste = sorted([f for f in listdir(tested_path) if isfile(join(tested_path, f)) and f in expected_liste and not f.startswith(".")])

        print(expected_liste == files_liste) # True for PASS, False for FAIL
        # END OF SCRIPT

    else:
        print("Usage: files_presence.py <expected_file>")
