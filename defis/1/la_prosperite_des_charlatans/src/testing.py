#!/usr/bin/env python3
import json, os, sys
from os import listdir
from os.path import isfile, join
from subprocess import *
from termcolor import colored
from TestReport import TestReport
from ThreadExec import ThreadExec

with open("data/config.json") as data:
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

def createDirLogs(target):
    result_file = "results/" + target["name"]
    if not os.path.exists(result_file):
        os.makedirs(result_file)

def executeTestList(test_list, test_report=None, wait=True):
    running_tests = []

    for test in test_list:
        TEST_PATH = config["settings"]["test_dir"] + test + "/"

        threadExec = ThreadExec(TEST_PATH, target, test_report)
        running_tests.append(threadExec)
        threadExec.start()

        if wait : threadExec.join()

    for test in running_tests:
        test.join()

if __name__ == "__main__":
    grp_name, test_number = parseArgs(sys.argv[1::])
    targets = getTargets(grp_name)

    INIT_TEST = ["make_target"]
    CLEANUP_TEST = ["make_mrproper"]

    TEST_LIST = getTestList()
    nb_tests_total = len(TEST_LIST) - len(CLEANUP_TEST)
    TEST_LIST = [t for t in TEST_LIST if t not in INIT_TEST and t not in CLEANUP_TEST]

    if test_number != None:
        TEST_LIST = [TEST_LIST[int(test_number)]]

    for target in targets:
        createDirLogs(target)
        test_report = TestReport(target, "results/" + target["name"], "rapport")

        print(("-"*64) + "\n # " + target["name"].upper() + "\n" + ("-"*64))

        executeTestList(CLEANUP_TEST, test_report)
        executeTestList(INIT_TEST, test_report)
        executeTestList(TEST_LIST, test_report, wait=False)
        executeTestList(CLEANUP_TEST)

        test_report.saveReportDatabase()
        test_report.saveReportHtml()

        total_exec_time = test_report.get_total_exec_time()[1]

        colortotal = "green" if test_report.NB_PASSED == nb_tests_total else "yellow"
        print (colored("\n [{}/{}] {}s => {:0.1f}/20 avec félicitations du jury\n".format(test_report.NB_PASSED, nb_tests_total, total_exec_time, (test_report.NB_PASSED / nb_tests_total) * 20 ), colortotal))
