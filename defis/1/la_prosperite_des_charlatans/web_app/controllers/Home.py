import cherrypy
from controllers.Base import BaseController
import os

class HomeController(BaseController):
    @cherrypy.expose
    def index(self):

        lasts_passed = self.get_lasts_passed()
        lasts_commit_best_group = self.get_lasts_commit_best_group()

        template_args = {
        "last_group_test":lasts_passed,
        "last_commit_best_group":lasts_commit_best_group
        }

        return self.render_template(template_args)

    def get_lasts_passed(self):

        qry = "SELECT substr(sha, -8) as commit_sha, \
            date_test, \
            name_group, \
            result, \
            sum(case when result='PASS' then 1 else 0 end), \
            COUNT(DISTINCT name_test) \
            FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
            GROUP BY name_group, commit_sha ORDER BY date_test DESC;".format()

        return self.executeQuery(qry)
