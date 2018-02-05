#!/usr/bin/env python3
import json, os, sys, threading
from subprocess import *
from os.path import isfile, join
from termcolor import colored

class ThreadExec(threading.Thread):
    def __init__(self, test_path, target, test_report=None):
        threading.Thread.__init__(self)
        self.test_path = test_path
        self.target = target
        self.test_report = test_report
        self.out = ["False", ""]
        self.result = "FAIL"

        self.TEST_INFO = self.getTestInfos(test_path)
        self.EXEC = self.TEST_INFO["exec"][0]
        self.TEST_FILE = join(test_path, self.TEST_INFO["exec"][1])

        self.createDirOutput(test_path)

    def run(self):

        try:
            self.out = self.executeTest(self.EXEC, self.TEST_FILE, self.target)
            self.out = self.out.split("#")
        except Exception as e:
            pass

        self.result = "PASS" if self.out[0] == "True" else "FAIL"

        if self.test_report != None:
            self.test_report.addLog(self)
            self.printResult()


    def executeTest(self, executable, test_file, target):
        try:
            cmd = [executable, test_file, target["path"]]
            out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
        except:
            out = "False"
        return out.strip("\n")

    def getTestInfos(self, test_path):
        with open(join(test_path, "infos.json")) as data:
            test_data = json.load(data)
            TEST_INFO = test_data["infos"]
        return TEST_INFO


    def createDirOutput(self, test_path):
        output_path = test_path + "output/"
        if not os.path.exists(output_path):
            os.makedirs(output_path)

    def printResult(self):
        print(colored(" [{}] {}".format(self.result, self.TEST_INFO["name"]), 'green' if self.out[0] == "True" else 'red'))
