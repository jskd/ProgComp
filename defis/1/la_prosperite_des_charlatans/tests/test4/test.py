#!/usr/bin/env python3
import json, os, sys, re, filecmp
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
    EXEC_PATH = settings["executable"]

if __name__ == "__main__":
    if len(sys.argv) >= 1:
        test_path = os.path.dirname(sys.argv[0]) + "/"
        group_path = sys.argv[1] + "/"

        # INIT DATA
        with open(test_path + INFOS_PATH, 'r') as data:
            test_data = json.load(data)
            test_info = test_data["infos"]
            test_timeout = int(test_info["timeout"])


        # BEGIN OF SCRIPT
        test_result = True
        try:
            input_data = os.path.realpath(test_path + "input/data.csv")
            input_user = os.path.realpath(test_path + "input/user.txt")
            view0_output = os.path.realpath(test_path + "output/view0.csv")
            changes_output = os.path.realpath(test_path + "output/changes.txt")

            # CLEAN OUTPUT
            if os.path.isdir(view0_output):
                os.rmdir(view0_output)
            else:
                if os.path.isfile(view0_output):
                    os.remove(view0_output)

            if os.path.isdir(changes_output):
                os.rmdir(changes_output)
            else:
                if os.path.isfile(changes_output):
                    os.remove(changes_output)

            cmd = [group_path + EXEC_PATH, input_data, input_user, view0_output, changes_output]
            out = check_output(cmd, stderr=STDOUT, timeout=test_timeout).decode("utf-8")

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

            if test_result:
                out = "(changes.txt == changes_expected.txt) AND (view0.csv == view0_expected.csv)"
            else:
                out = "(changes.txt != changes_expected.txt) OR (view0.csv != view0_expected.csv)"


        except Exception as e:
            out = str(e)
            test_result = False

        # True for PASS, False for FAIL
        print("{}#{}".format(test_result, out))
        # END OF SCRIPT
    else:
        print("Usage: files_presence.py <expected_file>")
