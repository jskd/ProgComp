import cherrypy, sqlite3
from controllers.Base import BaseController
import os
import datetime

class GraphicsController(BaseController):
    @cherrypy.expose
    def index(self):
        all_sha = self.all_sha_commit()
        all_test = self.all_name_test()
        all_group = self.all_name_group()

        groups = []


        for (group,) in all_group:
          groups.append( group)


        for (name_test,) in all_test:

          print("#date sha " + ' '.join(groups) )
          for (sha,) in all_sha:
            times = []
            for group in groups:
              times.append( self.get_time(name_test, sha, group) )
            print(self.get_date_of_sha_commit(sha)+ ' ' + sha + ' ' + ' '.join(times) )

        template_args = {
        "data_set":all_sha
        }
        return self.render_template(template_args)

    def all_sha_commit(self):
        qry = "SELECT distinct sha \
            FROM logs \
            ORDER BY date_test;".format()
        return self.executeQuery(qry)


    def all_name_test(self):
        qry = "SELECT distinct name_test \
            FROM logs \
            ORDER BY name_test;".format()
        return self.executeQuery(qry)

    def all_name_group(self):
        qry = "SELECT distinct name_group \
            FROM logs \
            ORDER BY name_group;".format()
        return self.executeQuery(qry)

    def get_time(self, name_test, sha, name_group):
        qry = "SELECT bench_time \
            FROM logs \
            WHERE name_test = '{}' AND sha = '{}' AND name_group = '{}' \
            ORDER BY name_group \
            LIMIT 1;".format(name_test.replace("'","''"), sha, name_group)
        try:
          return str(self.executeQuery(qry)[0][0])
        except:
          return "NaN"

    def get_date_of_sha_commit(self, sha):
        qry = "SELECT date_test \
            FROM logs \
            WHERE sha = '{}' \
            ORDER BY date_test \
            LIMIT 1;".format(sha)
        try:
          return str(datetime.datetime.strptime(self.executeQuery(qry)[0][0], "%Y-%m-%d %H:%M:%S").strftime("%d/%m/%Y"))
        except:
          return "NaN"