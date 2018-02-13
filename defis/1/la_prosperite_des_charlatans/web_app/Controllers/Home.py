import cherrypy
from Controllers.Base import BaseController
import os


class HomeController(BaseController):
    @cherrypy.expose
    def index(self):
        return self.render_template({"pwd":os.getcwd()})
