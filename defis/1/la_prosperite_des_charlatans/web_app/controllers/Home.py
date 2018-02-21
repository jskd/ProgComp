import cherrypy
from controllers.Base import BaseController
import os

class HomeController(BaseController):
    @cherrypy.expose
    def index(self):

        group_stats = self.get_group_stats()
        lasts_passed = self.get_lasts_passed()

        template_args = {
        "group_stats":group_stats,
        "last_group_test":lasts_passed
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
