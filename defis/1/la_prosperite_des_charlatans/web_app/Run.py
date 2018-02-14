from controllers.Home import HomeController
import cherrypy, os
from cherrypy import tools

def error_page_404(status, message, traceback, version):
    return "404 Error!"

def start_server():
    site_config_path = 'web_app/site.conf'

    cherrypy.tree.mount(HomeController(), '/', site_config_path)
    cherrypy.config.update({'error_page.404': error_page_404})
    cherrypy.config.update('web_app/server.conf')
    cherrypy.engine.start()

if __name__ == '__main__':
    start_server()
