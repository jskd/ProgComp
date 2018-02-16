import cherrypy, sqlite3
from controllers.Base import BaseController
import os

class StatisticsController(BaseController):
    @cherrypy.expose
    def index(self):
        lasts_commit_best_group = self.get_lasts_commit_best_group()
        group_stats = self.get_group_stats()

        template_args = {
        "last_commit_best_group":lasts_commit_best_group,
        "group_stats":group_stats
        }
        return self.render_template(template_args)

    def get_group_stats(self):
        #qry = "SELECT DISTINCT name_group FROM logs; ".format()

        qry = "SELECT sha as commit_sha, \
            date_test, \
            name_group, \
            result, \
            sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
            COUNT(DISTINCT name_test) as total_test \
            FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
            WHERE commit_sha = (SELECT sha FROM logs ORDER BY date_test DESC LIMIT 1)\
            GROUP BY commit_sha, name_group \
            ORDER BY nb_passed DESC;".format()

        return self.executeQuery(qry)
