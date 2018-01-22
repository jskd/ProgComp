#!/usr/bin/env python3

import json
import os,sys
from subprocess import *
from collections import namedtuple


PASS = "PASS"
FAIL = "FAIL"
ResultTest = namedtuple("ResultTest", "sucess name desc change stdout view0")

def make_target(target):
    print("[MAKE] " + target["name"])
    print("[FAIL] -> A implement")

def testSuit(target, files_config):
    make_target(target)
    index = 0
    for test_dir in os.listdir( files_config["test-dir"] ):
        test_path = files_config["test-dir"]+"/"+test_dir
        result= unitTest(test_path, target, files_config)
        #result= unitTest(target, files_config)
        print("#{} [{}] {} {}".format(index ,result.sucess, result.name, result.desc))
        index += 1

def unitTest(test_dir, target, files_config):

    with open(test_dir + "/config.json") as data:
        test_config = json.load(data)

    output_dir = test_dir + files_config["output-dir"] + "/" +target["name"]


    view_out = output_dir + files_config["output-view0"]
    view_exp = test_dir + files_config["expected-view0"]

    diff_view  = Popen(["diff", view_out, view_exp ], stdout=PIPE).communicate()[0].decode("utf-8")

    #output_dir = files_config["output-dir"] + "/"+target["name"]

    #print(expected_dir)
    #print(output_dir + files_config['output-view0'])

    #print(test_dir + files_config["output-dir"])

    #

    #print(output_dir + files_config["output-view0"])
    #print(files_config['expected-view0'])
    #diff_change= Popen(["diff", CHANGE_OUT, CHANGE_EXPECTED], stdout=PIPE).communicate()[0].decode("utf-8")

    #"": "output/changes.txt",
    #"": "output/view0.txt",
    #"output-stdio": "output/stdio.txt",


    return ResultTest(sucess=FAIL, name=test_config["name"],
        desc=test_config["desc"], change=PASS, stdout=PASS, view0=PASS)




'''
def uniteTest(test_dir, group_name):

    init_group_dir(test_dir, group_name);

    GROUP_OUTPUT    = get_group_dir(test_dir, group_name)
    VIEW_OUT        = GROUP_OUTPUT + "/view0.csv"
    CHANGE_OUT      = GROUP_OUTPUT + "/changes.txt"

    cmd= ["echo", DATA_IN, USER_IN, VIEW_OUT, CHANGE_OUT ]


    out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")


    if diff_view != "" :
        print("[FAIL-VIEW] ")

    if diff_change != "" :
        print("[FAIL-CHANGE] ")
'''

if __name__ == "__main__":


    with open("config.json") as data:
        config = json.load(data)

    for target in config["targets"]:
        testSuit(target, config["files-config"])


    #uniteTest("tests/0", "toto")
