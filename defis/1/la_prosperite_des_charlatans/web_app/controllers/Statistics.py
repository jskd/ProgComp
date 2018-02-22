import cherrypy, sqlite3
from controllers.Base import BaseController
import os

class StatisticsController(BaseController):
    @cherrypy.expose
    def index(self):
        group_stats = self.get_group_stats()

        template_args = {
        "group_stats":group_stats
        }
        return self.render_template(template_args)

    @cherrypy.expose
    def history(self, group):
        group_stats = self.get_group_stats()
        group_history = self.get_history(group)

        template_args = {
        "group_stats":group_stats,
        "group_history":group_history
        }
        return self.render_template(template_args)

    @cherrypy.expose
    def detail_list(self, **args):
        group_stats = self.get_group_stats()
        detail_list_test = self.get_detail_list_test(args["sha"], args["group"])

        template_args = {
        "group_stats":group_stats,
        "detail_list_test":detail_list_test
        }
        return self.render_template(template_args)

    @cherrypy.expose
    def detail(self, id):
        group_stats = self.get_group_stats()
        detail_test = self.get_detail_test(id)

        template_args = {
        "group_stats":group_stats,
        "detail_test":detail_test
        }
        return self.render_template(template_args)

    def get_history(self, group):

        qry = "SELECT sha, \
            date_test, \
            name_group, \
            sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
            COUNT(DISTINCT name_test) as total_test \
            FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
            WHERE name_group = '{}' \
            GROUP BY sha, name_group \
            ORDER BY date_test DESC".format(group)

        return self.executeQuery(qry)

    def get_detail_list_test(self, sha, group):

        qry = "SELECT * \
            FROM logs \
            WHERE sha = '{}' AND name_group = '{}' \
            ORDER BY date_test".format(sha, group)

        test_list = self.executeQuery(qry)

        for i, test in enumerate(test_list):
            test_name = test[4]

            qry = "SELECT sha \
                FROM logs \
                WHERE name_group = \"{}\" and  name_test = \"{}\" and result='PASS' \
                ORDER BY date_test DESC LIMIT 1;".format(group, test_name)

            last_valid_sha = self.executeQuery(qry)

            if last_valid_sha != None and len(last_valid_sha) > 0:
                test_list[i] += last_valid_sha[0]
            else:
                test_list[i] += ('',)

        return test_list

    def get_detail_test(self, id):

        qry = "SELECT * \
            FROM logs \
            WHERE id = '{}';".format(id)

        test_detail = self.executeQuery(qry)[0]

        group_name = test_detail[3]
        test_name = test_detail[4]

        qry = "SELECT sha \
            FROM logs \
            WHERE name_group = \"{}\" and  name_test = \"{}\" and result='PASS' \
            ORDER BY date_test DESC LIMIT 1;".format(group_name, test_name)

        last_valid_sha = self.executeQuery(qry)

        if last_valid_sha != None and len(last_valid_sha) > 0:
            test_detail += last_valid_sha[0]
        else:
            test_detail += ('',)

        return test_detail
