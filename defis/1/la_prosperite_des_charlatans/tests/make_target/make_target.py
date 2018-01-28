#!/usr/bin/env python3
import json, os, sys, re
from subprocess import *
from os import listdir
from os.path import isfile, join

# OPEN CONSTANTS IN config.json
with open("config.json") as config_data:
    config = json.load(config_data)

    BIN_PATH = config["settings"][1]["bin_dir"]
    MAKEFILE_PATH = config["settings"][2]["makefile"]
    EXPECTED_PATH = config["settings"][3]["expected_dir"]
    INFOS_PATH = config["settings"][4]["infos_json"]

if __name__ == "__main__":
    if len(sys.argv) >= 1:
        test_path = os.path.dirname(sys.argv[0]) + "/"
        group_path = sys.argv[1] + "/"

        # INIT DATA
        with open(test_path + INFOS_PATH, 'r') as data:
            test_data = json.load(data)
            test_info = test_data["infos"][0]

        # BEGIN OF SCRIPT
        test_result = True
        try:
            #if not os.path.isdir(group_path + BIN_PATH):
            #    raise Exception("Error: bin/ does not exist.")

            with open(group_path + MAKEFILE_PATH, "r"): pass

            cmd = ["make", "--directory=" + group_path]
            out = check_output(cmd, stderr=STDOUT, timeout=60).decode("utf-8")
            """
            matchObj = re.match(r"make\[1\]*", out)
            if matchObj:
                test_result = False
            """
        except:
            out = "False"
            test_result = False

        # True for PASS, False for FAIL
        print(test_result)
        # END OF SCRIPT
    else:
        print("Usage: files_presence.py <expected_file>")
