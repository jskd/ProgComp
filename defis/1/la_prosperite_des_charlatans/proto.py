#!/usr/bin/env python3

import json
import os,sys
from subprocess import *

if __name__ == "__main__":

    with open("target.json") as data:
        data_targets = json.load(data)

    with open("tests.json") as data:
        data_tests = json.load(data)


    for target in data_targets["target"]:
        print(" # Current project :", target["name"])

        for i, test in enumerate(data_tests["tests"]):

            cmd = test["exec"]

            if cmd == "" :
                # Here, tests without exec files
                # out = anotherFunction()
                out = ""
            else:
                out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")

            if test["expected"] != "":
                with open(test["expected"], 'r') as f:
                    result = "PASS" if f.read() == out else "FAIL"
            else:
                # Here, tests without expected files
                # result = anotherExpectedThings
                result = "FAIL"

            print(" [{}] -> [{}] {} {}".format(i, result, test["name"], test["expected"]))

    print ("\n >> RESULT : [{}/{}] -> NOTE: 20/20 avec félicitations du jury".format(0, len(data_tests["tests"])))
    print ()
