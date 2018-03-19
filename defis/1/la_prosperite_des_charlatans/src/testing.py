#!/usr/bin/env python3
import json, os, sys
from os import listdir
from os.path import isfile, join
from subprocess import *
from termcolor import colored
from TestReport import TestReport
#from ThreadExec import ThreadExec
from TaskTracker import TaskTracker

with open("data/config.json") as data:
    config = json.load(data)

def parseArgs(args):
    """ Parse arguments from line """
    defis_number = None
    grp_name = None
    test_number = None

    for arg in args:
        splited = arg.split("=")
        if splited[0] == "defis":
            defis_number = splited[1]
            if defis_number == "" : defis_number = 1
        elif splited[0] == "grp":
            grp_name = splited[1]
        elif splited[0] == "number":
            try: test_number = int(splited[1])
            except: pass

    return defis_number, grp_name, test_number

def getTargets(defis_number, grp_name = None):
    """ Get groups in config.json """
    targets = config["target"]
    for target in targets:
        target["path"] = target["path"].replace("${DEFIS_NUMBER}", str(defis_number))

        if target["name"] == grp_name:
            targets = [target]
            break
    return targets

def getTestList():
    """ Get all tests """
    TEST_LIST = []

    for test_dir in os.listdir("tests/"):
        if os.path.isdir("tests/" + test_dir):
            TEST_LIST += [os.path.basename(test_dir)]

    return TEST_LIST

def createDirLogs(target):
    """ Create output dir for reports """
    result_file = "results/" + target["name"]
    if not os.path.exists(result_file):
        os.makedirs(result_file)

def executeTestList(test_list, test_report=None, wait=True):
    """ Execute a test list """

    running_tests = []

    for test in test_list:
        TEST_PATH = config["settings"]["test_dir"] + test + "/"

        taskTracker = TaskTracker(TEST_PATH, target, test_report)
        running_tests.append(taskTracker)
        taskTracker.start()

        if wait : taskTracker.join()

    for test in running_tests:
        test.join()

if __name__ == "__main__":

    # INITIALIZATION OF CONTEXT
    #------------------------------------------------------
    defis_number, grp_name, test_number = parseArgs(sys.argv[1::])
    targets = getTargets(defis_number, grp_name)

    INIT_TEST = ["make_target"]
    CLEANUP_TEST = ["make_mrproper"]

    TEST_LIST = getTestList()
    TEST_LIST = [t for t in TEST_LIST if t not in INIT_TEST and t not in CLEANUP_TEST]

    if test_number != None:
        TEST_LIST = [TEST_LIST[int(test_number)]]
    #------------------------------------------------------


    # EXECUTION OF TESTS
    #------------------------------------------------------
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

        colortotal = "green" if test_report.NB_PASSED == test_report.NB_TEST else "yellow"
        print (colored("\n [{}/{}] {}s => {:0.1f}/20 avec félicitations du jury\n".format(test_report.NB_PASSED, test_report.NB_TEST, total_exec_time, (test_report.NB_PASSED / test_report.NB_TEST) * 20 ), colortotal))
