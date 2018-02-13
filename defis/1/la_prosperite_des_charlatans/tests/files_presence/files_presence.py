#!/usr/bin/env python3
import json, os, sys, re
from subprocess import *
from os import listdir
from os.path import isfile, join

# OPEN CONSTANTS IN config.json
with open("data/config.json") as config_data:
    config = json.load(config_data)
    settings = config["settings"]

    BIN_PATH = settings["bin_dir"]
    MAKEFILE_PATH = settings["makefile"]
    EXPECTED_PATH = settings["expected_dir"]
    INFOS_PATH = settings["infos_json"]
    RESULT_PATH = settings["result_dir"]

if __name__ == "__main__":
    if len(sys.argv) >= 1:
        test_path = os.path.dirname(sys.argv[0]) + "/"
        group_path = sys.argv[1] + "/"

        # INIT DATA
        with open(test_path + INFOS_PATH, 'r') as data:
            test_data = json.load(data)
            test_info = test_data["infos"]

        # BEGIN OF SCRIPT
        expected_path = join(test_path, EXPECTED_PATH + test_info["expected"])
        with open(expected_path, 'r') as expected_files:
            expected_liste = sorted(expected_files.read().strip("\n").split("\n"))

        files_liste = []
        for path, subdirs, files in os.walk(group_path):
            for name in files:
                if not name.startswith(".") and name in expected_liste: files_liste.append(name)
        files_liste = list(set(files_liste))

        # True for PASS, False for FAIL
        print(sorted(expected_liste) == sorted(files_liste))
        # END OF SCRIPT
    else:
        print("Usage: files_presence.py <expected_file>")
