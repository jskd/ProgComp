#!/usr/bin/env python3

import json
import os,sys
from os import listdir
from os.path import isfile, join
from subprocess import *

TEST_DIR = "tests/"
OUT_EXPECTED_DIR = "output-expected/"

if __name__ == "__main__":

    with open("target.json") as data:
        data_targets = json.load(data)

    with open("tests.json") as data:
        data_tests = json.load(data)

    tests_folders = [f for f in listdir(TEST_DIR) if not isfile(join(TEST_DIR, f)) and not f.startswith(".")]

    for target in data_targets["target"]:
        print(" # Current project :", target["name"])
        print("-"*32)

        nb_passed = 0

        for i, test in enumerate(tests_folders):

            infos_path = join(TEST_DIR + test + "/", "infos.json")

            with open(infos_path) as data:
                test_data = json.load(data)

            test_info = test_data["infos"][0]
            cmd = test_info["exec"]

            try:
                cmd = [cmd[0], join(TEST_DIR + test + "/", cmd[1]), target["path"], infos_path]
            except:
                cmd = test_info["exec"]

            out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")

            if out.strip("\n") == "True" :
                result = "PASS"
                nb_passed += 1
            else:
                result = "FAIL"


            print(" [{}] -> [{}] {} {}".format(i, result, test_info["name"], test_info["expected"]))

        print (" TOTAL : [{}/{}] -> NOTE: {}/20 avec félicitations du jury".format(nb_passed, len(tests_folders), (nb_passed / len(tests_folders)) * 20 ))
        print ()
