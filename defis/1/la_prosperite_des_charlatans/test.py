#!/usr/bin/env python3

import json
import os,sys
from subprocess import *

def uniteTest(test_dir, group_name):

    init_group_dir(test_dir, group_name);

    GROUP_OUTPUT    = get_group_dir(test_dir, group_name)
    VIEW_OUT        = GROUP_OUTPUT + "/view0.csv"
    CHANGE_OUT      = GROUP_OUTPUT + "/changes.txt"

    cmd= ["echo", DATA_IN, USER_IN, VIEW_OUT, CHANGE_OUT ]


    out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")

    diff_view  = Popen(["diff", VIEW_OUT  , VIEW_EXPECTED], stdout=PIPE).communicate()[0].decode("utf-8")
    diff_change= Popen(["diff", CHANGE_OUT, CHANGE_EXPECTED], stdout=PIPE).communicate()[0].decode("utf-8")

    if diff_view != "" :
        print("[FAIL-VIEW] ")

    if diff_change != "" :
        print("[FAIL-CHANGE] ")


if __name__ == "__main__":

    with open("target.json") as data:
        data_targets = json.load(data)

    with open("tests.json") as data:
        data_tests = json.load(data)

    default_file_config = data_tests["default-file-config"]

    print(default_config)



    #uniteTest("tests/0", "toto")
