#!/usr/bin/env python3
import json, os, sys
from subprocess import *

class TestReport():
    def __init__(self, target, report_path, report_name):
        self.report_path = report_path
        self.report_name = report_name
        self.logs = []
        self.target = target
        self.NB_PASSED = 0

    def addLog(self, threadExec):
        log = {
            "test_info" : threadExec.TEST_INFO,
            "result" : threadExec.result,
            "output" : threadExec.out[1] if len(threadExec.out) > 1 else ""
        }

        if log["result"] == "PASS": self.NB_PASSED += 1
        self.logs.append(log)

    def saveReport(self):
        self.report_name += ".html"

        try:
            cmd = ["git", "rev-parse", "HEAD"]
            sha = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
        except: sha = ""

        try:
            cmd = ["cp", "results/form.html", self.report_path + "/" + self.report_name]
            cmd_out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
        except: pass

        with open(self.report_path + "/" + self.report_name, "r") as rapport:
            rapport_string = rapport.read()


        rapport_string = rapport_string.replace("form.css", "../form.css")
        rapport_string = rapport_string.replace("${GROUP}", self.target["name"])
        rapport_string = rapport_string.replace("${SHA}", sha)

        test_list = ""
        for log in self.logs:
            current_log = " \
            <div class='block_test'> \n\
            <div class='title_test'> \n\
            <p><font color='#2dae5b'>${TEST_NAME}</font></p> \n\
            <div class='liste_actions'> \n\
            <p><font color='#2dae5b'>[${RESULT}]</font></p> \n\
            <a class='result_icon_pass' href='' alt='Result' title='Result'></a> \n\
            </div> \n\
            </div> \n\
            <h4>OUTPUT :</h4> \n\
            <p> \n\
            ${OUTPUT} \n\
            </p> \n\
            </div> \n"

            current_log = current_log.replace("${TEST_NAME}", log["test_info"]["name"])
            current_log = current_log.replace("${RESULT}", log["result"])

            output = log["output"].replace("\n", "<br/>")
            current_log = current_log.replace("${OUTPUT}", output)

            if log["result"] == "FAIL":
                current_log = current_log.replace("#2dae5b", "red")
                current_log = current_log.replace("result_icon_pass", "result_icon_fail")

            test_list += current_log

        rapport_string = rapport_string.replace("${TEST_LIST}", test_list)

        with open(self.report_path + "/" + self.report_name, "w") as rapport:
            rapport.write(rapport_string)
