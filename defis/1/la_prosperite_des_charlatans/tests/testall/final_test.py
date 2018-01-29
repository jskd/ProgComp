#!/usr/bin/env python3
import json, os, sys, re, filecmp
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
    EXEC_PATH = config["settings"][7]["executable"]

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
            view0_output = test_path + "output/view0.csv"
            changes_output = test_path + "output/changes.txt"

            # CLEAN OUTPUT
            os.remove(view0_output)
            os.remove(changes_output)

            cmd = [group_path + EXEC_PATH, test_path + "input/data.csv", test_path + "input/user.txt", view0_output, changes_output]
            out = check_output(cmd, stderr=STDOUT, timeout=30).decode("utf-8")

            changes_expected = test_path + EXPECTED_PATH + "changes_expected.txt"
            view0_expected = test_path + EXPECTED_PATH + "view0_expected.csv"


            with open(changes_expected, 'r') as expected, open(changes_output, 'r') as output:
                expected_strip = expected.read().strip("\n")
                output_strip = output.read().strip("\n")

                test_result = (output_strip == expected_strip)

            with open(view0_expected, 'r') as expected, open(view0_output, 'r') as output:
                expected_strip = expected.read().strip("\n")
                output_strip = output.read().strip("\n")

                test_result = test_result and (output_strip == expected_strip)

        except:
            test_result = False

        # True for PASS, False for FAIL
        print(test_result)
        # END OF SCRIPT
    else:
        print("Usage: files_presence.py <expected_file>")