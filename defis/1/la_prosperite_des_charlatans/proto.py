#!/usr/bin/env python3

import json
import os,sys
from os import listdir
from os.path import isfile, join
from subprocess import *
from termcolor import colored

with open("config.json") as data:
    config = json.load(data)

TEST_DIR = config["settings"][0]["test_dir"]
TEST_LIST = sorted([f for f in listdir(TEST_DIR) if not isfile(join(TEST_DIR, f)) and not f.startswith(".")])

if __name__ == "__main__":
    targets = config["target"]

    if len(sys.argv) > 1:
        grp_name = sys.argv[1]

        for target in targets:
            if target["name"] == grp_name:
                targets = [target]
                break

    for target in targets:

        NB_PASSED = 0
        print("-"*64)
        print(" #", target["name"].upper())
        print("-"*64)

        for i, test in enumerate(TEST_LIST):
            infos_path = join(TEST_DIR + test + "/", "infos.json")

            with open(infos_path) as data:
                test_data = json.load(data)

            test_info = test_data["infos"][0]
            cmd = test_info["exec"]
            exec_file = join(TEST_DIR + test + "/", cmd[1])

            try:
                cmd = [cmd[0], exec_file, target["path"]]
                out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
            except:
                out = "False"

            if out.strip("\n") == "True" :
                result = "PASS"
                colorprint = 'green'
                NB_PASSED += 1
            else:
                result = "FAIL"
                colorprint = 'red'

            print(colored(" [{}] -> [{}] {} {}".format(i, result, test_info["name"], test_info["expected"]), colorprint))

        colortotal = "green" if NB_PASSED == len(TEST_LIST) else "yellow"

        print (colored("\n [{}/{}] => {}/20 avec félicitations du jury\n".format(NB_PASSED, len(TEST_LIST), (NB_PASSED / len(TEST_LIST)) * 20 ), colortotal))
