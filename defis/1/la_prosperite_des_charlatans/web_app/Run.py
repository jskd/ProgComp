from controllers.Home import HomeController
from controllers.Statistics import StatisticsController
from controllers.Graphics import GraphicsController
import cherrypy, os, re
from cherrypy import tools

def error_page_404(status, message, traceback, version):
    return "404 Error!"

def start_server():
    site_base_config_path = 'web_app/site_base.conf'
    site_config_path = 'web_app/site.conf'

    with open(site_base_config_path, 'r') as f:
        config = f.read()
        config = config.replace("${STATIC_DIR}", "\"" + os.getcwd() + "/web_app" + "\"")

    with open(site_config_path, 'w') as f:
        f.write(config)

    cherrypy.tree.mount(HomeController(), '/', site_config_path)
    cherrypy.tree.mount(StatisticsController(), '/statistics/', site_config_path)
    cherrypy.tree.mount(GraphicsController(), '/graphics/', site_config_path)
    cherrypy.config.update({'error_page.404': error_page_404})
    cherrypy.config.update('web_app/server.conf')
    cherrypy.engine.start()

if __name__ == '__main__':
    start_server()
