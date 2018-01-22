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
        print(True) # True for PASS, False for FAIL
        # END OF SCRIPT

    else:
        print("Usage: files_presence.py <expected_file>")
