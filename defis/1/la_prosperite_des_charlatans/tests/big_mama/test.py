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
        size_bigmama = sys.argv[2]

        # INIT DATA
        with open(test_path + INFOS_PATH, 'r') as data:
            test_data = json.load(data)
            test_info = test_data["infos"]
            test_timeout = int(test_info["timeout"])

        # BEGIN OF SCRIPT
        test_result = True
        try:
            view0_output = test_path + "output/view0_{}.csv".format(size_bigmama)
            changes_output = test_path + "output/changes_{}.txt".format(size_bigmama)

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

            intput_file = "input/bigmama_{}.csv".format(size_bigmama)

            if not os.path.isfile(test_path + intput_file):
                raise Exception("{} doesn't exists.".format(intput_file))

            cmd = [group_path + EXEC_PATH, test_path + intput_file, test_path + "input/user.txt", view0_output, changes_output]
            out = check_output(cmd, stderr=STDOUT, timeout=test_timeout).decode("utf-8")

        except Exception as e:
            out = str(e)
            test_result = False

        # True for PASS, False for FAIL
        print("{}#{}".format(test_result, out.strip("\n")))
        # END OF SCRIPT
    else:
        print("Usage: files_presence.py <expected_file>")
