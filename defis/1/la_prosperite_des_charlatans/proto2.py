#!/usr/bin/env python3

import json
import os,sys
from subprocess import *

if __name__ == "__main__":

    with open("target.json") as data:
        data_targets = json.load(data)

    with open("tests2.json") as data:
        data_tests = json.load(data)


    for target in data_targets["target"]:
        print(" # Current project :", target["name"])
        print("-"*32)

        nb_passed = 0
        Popen(["ws", "tests/data.csv", "tests/user.txt",target["path"]+  "view0.csv", target["path"]+ "changes.txt"], stdout=PIPE).communicate()[0].decode("utf-8")

        for i, test in enumerate(data_tests["tests"]):
            test_file = target["path"] +"/"+test["expected"]
            cmd = ["diff",test_file ,test["file"] ]
            print(cmd)
            out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
            """
            if cmd == "" :
                # Here, tests without exec files
                # out = anotherFunction()
                out = "0 0 5\n1 0 10\n1 1 =#(0, 0, 1, 0, 5)"
            else:
                

            if test["expected"] != "":
                with open(test["expected"], 'r') as f:
                    result = "PASS" if f.read().strip("\n") == out else "FAIL"
                    if result == "PASS" : nb_passed += 1

            else:
                # Here, tests without expected files
                # result = anotherExpectedThings
                result = "FAIL"

            print(" [{}] -> [{}] {} {}".format(i, result, test["name"], test["expected"]))
            """

        #print (" TOTAL : [{}/{}] -> NOTE: {}/20 avec félicitations du jury".format(nb_passed, len(data_tests["tests"]), (nb_passed / len(data_tests["tests"])) * 20 ))
        #print ()
