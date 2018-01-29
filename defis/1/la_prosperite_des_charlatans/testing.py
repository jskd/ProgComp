#!/usr/bin/env python3
import json, os, sys
from os import listdir
from os.path import isfile, join
from subprocess import *
from termcolor import colored

with open("config.json") as data:
    config = json.load(data)

def parseArgs(args):
    grp_name = None
    test_number = None
    for arg in args:
        splited = arg.split("=")
        if splited[0] == "grp":
            grp_name = splited[1]
        elif splited[0] == "number":
            try: test_number = int(splited[1])
            except: pass
    return grp_name, test_number

def getTargets(grp_name = None):
    targets = config["target"]
    for target in targets:
        if target["name"] == grp_name:
            targets = [target]
            break
    return targets

def getTestList():
    TEST_LIST = []
    with open("tests/_tests_order.txt") as tests_order:
        for line in tests_order.readlines():
            line = line.strip("\n")
            if os.path.isdir("tests/" + line):
                TEST_LIST += [line]
    return TEST_LIST

def getTestInfos(test_path):
    with open(join(test_path, "infos.json")) as data:
        test_data = json.load(data)
        TEST_INFO = test_data["infos"]
    return TEST_INFO

def createDirLogs(target):
    result_file = "results/" + target["name"]
    if not os.path.exists(result_file):
        os.makedirs(result_file)
    with open(result_file + "/rapport.txt","w"): pass

def executeTest(executable, test_file, target):
    try:
        cmd = [executable, test_file, target["path"]]
        out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
        #print(out)
    except:
        out = "False"
    return out.strip("\n")


if __name__ == "__main__":
    grp_name, test_number = parseArgs(sys.argv[1::])
    targets = getTargets(grp_name)
    TEST_LIST = getTestList()

    if test_number != None:
        TEST_LIST = [TEST_LIST[int(test_number)]]

    for target in targets:
        createDirLogs(target)
        NB_PASSED = 0

        print(("-"*64) + "\n # " + target["name"].upper() + "\n" + ("-"*64))

        for test in TEST_LIST:
            TEST_PATH = config["settings"]["test_dir"] + test + "/"
            TEST_INFO = getTestInfos(TEST_PATH)
            EXEC = TEST_INFO["exec"][0]
            TEST_FILE = join(TEST_PATH, TEST_INFO["exec"][1])

            out = executeTest(EXEC, TEST_FILE, target)

            if out == "True" :
                result = "PASS"
                colorprint = 'green'
                NB_PASSED += 1
            else:
                result = "FAIL"
                colorprint = 'red'

            print(colored(" [{}] {}".format(result, TEST_INFO["name"]), colorprint))

        colortotal = "green" if NB_PASSED == len(TEST_LIST) else "yellow"
        print (colored("\n [{}/{}] => {:0.1f}/20 avec félicitations du jury\n".format(NB_PASSED, len(TEST_LIST), (NB_PASSED / len(TEST_LIST)) * 20 ), colortotal))
