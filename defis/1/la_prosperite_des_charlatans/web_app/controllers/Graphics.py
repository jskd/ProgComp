import cherrypy, sqlite3
from controllers.Base import BaseController
import os

class GraphicsController(BaseController):
    @cherrypy.expose
    def index(self):
        lasts_commit_best_group = self.get_lasts_commit_best_group()

        template_args = {
        "last_commit_best_group":lasts_commit_best_group
        }
        return self.render_template(template_args)
