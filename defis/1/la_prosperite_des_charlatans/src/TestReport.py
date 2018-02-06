#!/usr/bin/env python3
import json, os, sys, datetime
from subprocess import *

class TestReport():
    def __init__(self, target, report_path, report_name):
        self.report_path = report_path
        self.report_name = report_name
        self.logs = []
        self.target = target
        self.NB_PASSED = 0

        try:
            cmd = ["git", "rev-parse", "HEAD"]
            self.commit = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
            self.commit = self.commit.strip("\n")
        except: self.commit = ""

    def addLog(self, threadExec):
        log = {
            "test_info" : threadExec.TEST_INFO,
            "result" : threadExec.result,
            "output" : threadExec.out[1] if len(threadExec.out) > 1 else "",
            "exec_time" : threadExec.execution_time,
            "commit" : self.commit
        }

        if log["result"] == "PASS": self.NB_PASSED += 1
        self.logs.append(log)

    def get_total_exec_time(self):
        total = 0.0
        for log in self.logs:
            total += log["exec_time"]

        seconds = int(total)
        m, s = divmod(seconds, 60)
        h, m = divmod(m, 60)

        return total, "{:02d}:{:02d}:{:02d}".format(h, m, s)

    def saveReportDatabase(self):
        now = datetime.datetime.now()
        with open("database.csv", "a") as database:
            for log in self.logs:

                database.write("{};{};{};{};{}\n".format(
                self.commit,
                now.strftime("%Y-%m-%d %H:%M"),
                self.target["name"],
                log["test_info"]["name"],
                log["result"]))


    def saveReportHtml(self):
        self.report_name += ".html"

        try:
            cmd = ["cp", "templates/form.html", self.report_path + "/" + self.report_name]
            cmd_out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")

            cmd = ["cp", "templates/form.css", self.report_path + "/"]
            cmd_out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")

            cmd = ["cp", "-r", "templates/img", self.report_path + "/"]
            cmd_out = Popen(cmd, stdout=PIPE).communicate()[0].decode("utf-8")
        except Exception as e:
            print(str(e))

        with open(self.report_path + "/" + self.report_name, "r") as rapport:
            rapport_string = rapport.read()


        #rapport_string = rapport_string.replace("form.css", "../form.css")
        #rapport_string = rapport_string.replace("img/", "../img/")
        rapport_string = rapport_string.replace("${GROUP}", self.target["name"])
        rapport_string = rapport_string.replace("${SHA}", self.commit)
        rapport_string = rapport_string.replace("${NB_PASSED}", str(self.NB_PASSED) + "/" + str(len(self.logs)))
        rapport_string = rapport_string.replace("${NB_FAILED}", str(len(self.logs) - self.NB_PASSED) + "/" + str(len(self.logs)))
        rapport_string = rapport_string.replace("${TOTAL_TIME}", self.get_total_exec_time()[1])

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
            <hr> \n\
            Executed in : <strong>${EXEC_TIME} sec.</strong> \n\
            </div> \n"

            current_log = current_log.replace("${TEST_NAME}", log["test_info"]["name"])
            current_log = current_log.replace("${RESULT}", log["result"])

            output = log["output"].replace("\n", "<br/>")
            current_log = current_log.replace("${OUTPUT}", output)
            current_log = current_log.replace("${EXEC_TIME}", "{:0.5f}".format(log["exec_time"]))

            if log["result"] == "FAIL":
                current_log = current_log.replace("#2dae5b", "red")
                current_log = current_log.replace("result_icon_pass", "result_icon_fail")

            test_list += current_log

        rapport_string = rapport_string.replace("${TEST_LIST}", test_list)

        with open(self.report_path + "/" + self.report_name, "w") as rapport:
            rapport.write(rapport_string)
