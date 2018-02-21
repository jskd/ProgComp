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
