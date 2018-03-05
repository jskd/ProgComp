#!/usr/bin/env python3
import json, os, sys, threading, time
from subprocess import *
from os.path import isfile, join
from termcolor import colored
from ThreadExec import ThreadExec

class TaskTracker(threading.Thread):
    """ TaskTracker keep informations of a test and run it. """

    def __init__(self, test_path, target, test_report=None):
        threading.Thread.__init__(self)
        self.test_path = test_path
        self.target = target
        self.test_report = test_report

        self.TEST_INFO = self.getTestInfos(test_path)
        self.test_base_name = self.TEST_INFO["name"]
        self.executable = self.TEST_INFO["exec"][0]
        self.test_file = join(test_path, self.TEST_INFO["exec"][1])

        self.createDirOutput(test_path)

    def run(self):
        """ Run N thread where N = size of args list in infos.json
            Wait all thread before return
        """
        running_tests = []

        if "args" in self.TEST_INFO :
            for args in self.TEST_INFO["args"]:

                infos = self.TEST_INFO.copy()
                infos["name"] = self.test_base_name.replace("${ARGS}", args)

                threadExec = ThreadExec(infos, self.test_file, self.executable, args, \
                    self.target, self.test_report)

                running_tests.append(threadExec)
                threadExec.start()
        else:
            threadExec = ThreadExec(self.TEST_INFO.copy(), self.test_file, self.executable, "", \
                self.target, self.test_report)

            running_tests.append(threadExec)
            threadExec.start()

        for test in running_tests:
            test.join()

    def getTestInfos(self, test_path):
        """ Return informations presents in infos.json """

        with open(join(test_path, "infos.json")) as data:
            test_data = json.load(data)
            TEST_INFO = test_data["infos"]
        return TEST_INFO


    def createDirOutput(self, test_path):
        output_path = test_path + "output/"
        if not os.path.exists(output_path):
            os.makedirs(output_path)
