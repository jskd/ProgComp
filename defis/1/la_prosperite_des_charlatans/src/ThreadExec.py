#!/usr/bin/env python3
import json, os, sys, threading, time
from subprocess import *
from os.path import isfile, join
from termcolor import colored

class ThreadExec(threading.Thread):
    """ ThreadExec execute one instance of test with arguments. """

    def __init__(self, test_info, test_file, executable, args, target, test_report):
        threading.Thread.__init__(self)
        self.test_info = test_info
        self.test_file = test_file
        self.executable = executable
        self.args = args
        self.target = target
        self.test_report = test_report
        self.out = ["False", ""]
        self.result = "FAIL"
        self.execution_time = 0.0

    def run(self):
        try:
            self.out = self.executeTest(self.executable, self.args, self.test_file, self.target)
            self.out = self.out.split("#")

        except Exception as e:
            print(e)

        self.result = "PASS" if self.out[0] == "True" else "FAIL"

        if self.test_report != None:
            self.test_report.addLog(self)
            self.printResult()


    def executeTest(self, executable, args, test_file, target):
        try:
            if args != '':
                cmd = [executable, test_file, target["path"], args]
            else:
                cmd = [executable, test_file, target["path"]]

            begin = time.time()
            out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
            self.execution_time = time.time() - begin
        except:
            out = "False"
        return out.strip("\n")


    def printResult(self):
        seconds = self.execution_time

        print(colored(" [{}] {:<8} : {} ".format(self.result, "{:0.3f}s".format(seconds) , self.test_info["name"]), 'green' if self.out[0] == "True" else 'red'))
