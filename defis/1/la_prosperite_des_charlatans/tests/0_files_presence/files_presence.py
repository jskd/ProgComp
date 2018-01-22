#!/usr/bin/env python3

import json
import os,sys
from subprocess import *
from os import listdir
import os.path
from os.path import isfile, join

if __name__ == "__main__":
    if len(sys.argv) >= 1:
        tested_path = sys.argv[1]
        infos_path = sys.argv[2]

        splited_infos = infos_path.split("/")
        TEST_DIR = "/".join(splited_infos[0:len(splited_infos)-1]) + "/"
        OUT_EXPECTED_DIR = "output-expected/"

        with open(infos_path, 'r') as data:
            test_data = json.load(data)

        test_info = test_data["infos"][0]
        expected_path = join(TEST_DIR, OUT_EXPECTED_DIR + test_info["expected"])

        with open(expected_path, 'r') as expected_files:
            expected_liste = sorted(expected_files.read().strip("\n").split("\n"))

        files_liste = sorted([f for f in listdir(tested_path) if isfile(join(tested_path, f)) and f in expected_liste and not f.startswith(".")])

        print(expected_liste == files_liste)

    else:
        print("Usage: files_presence.py <expected_file>")
