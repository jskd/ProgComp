import cherrypy, sqlite3
from controllers.Base import BaseController
import os

class StatisticsController(BaseController):
    @cherrypy.expose
    def index(self):
        lasts_commit_best_group = self.get_lasts_commit_best_group()
        liste_group = self.get_group_list()

        template_args = {
        "liste_group":liste_group,
        "last_commit_best_group":lasts_commit_best_group
        }
        return self.render_template(template_args)

    def get_group_list(self):
        qry = "SELECT DISTINCT name_group FROM logs; ".format()
        return self.executeQuery(qry)
