import cherrypy, sqlite3
from controllers.Base import BaseController
import os

DB_PATH = "data/database.db"

class HomeController(BaseController):
    @cherrypy.expose
    def index(self):

        lasts_passed = self.get_lasts_passed()
        lasts_commit_best_group = self.get_lasts_commit_best_group()

        template_args = {
        "last_group_test":lasts_passed,
        "last_commit_best_group":lasts_commit_best_group
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

    def get_lasts_commit_best_group(self):

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

    def executeQuery(self, query_string):
        result_qry = None
        try:
            conn = sqlite3.connect(DB_PATH)
            cursor = conn.cursor()

            cursor.execute(query_string)
            result_qry = cursor.fetchall()

        except Exception as e:
            print(str(e))
            conn.rollback()
        finally:
            cursor.close()
            conn.close()
            return result_qry
