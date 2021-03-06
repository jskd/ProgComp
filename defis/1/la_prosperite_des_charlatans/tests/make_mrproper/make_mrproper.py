#!/usr/bin/env python3
import json, os, sys, re
from subprocess import *
from os import listdir
from os.path import isfile, join

# OPEN CONSTANTS IN config.json
with open("data/config.json") as config_data:
    config = json.load(config_data)
    settings = config["settings"]


if __name__ == "__main__":
    if len(sys.argv) >= 1:
        test_path = os.path.dirname(sys.argv[0]) + "/"
        group_path = sys.argv[1] + "/"
        group_name = os.path.basename(os.path.normpath(group_path))

        # BEGIN OF SCRIPT
        test_result = True
        try:

            cmd = ["make", "--directory=" + group_path, "mrproper"]
            out = check_output(cmd, stderr=STDOUT, timeout=30).decode("utf-8")

        except Exception as e:
            out = str(e)
            test_result = False


        # True for PASS, False for FAIL
        print("{}#{}".format(test_result, out))
        # END OF SCRIPT
    else:
        print("Usage: files_presence.py <expected_file>")
