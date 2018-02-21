import cherrypy, sqlite3
from controllers.Base import BaseController
import os

class GraphicsController(BaseController):
    @cherrypy.expose
    def index(self):
        group_stats = self.get_group_stats()

        template_args = {
        "group_stats":group_stats
        }
        return self.render_template(template_args)
