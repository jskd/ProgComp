import cherrypy, sqlite3
from controllers.Base import BaseController
import os

DB_PATH = "data/database.db"

class HomeController(BaseController):
    @cherrypy.expose
    def index(self):
        lasts_passed = self.get_lasts_passed()
        return self.render_template({"last_group_test":lasts_passed})

    def get_lasts_passed(self):
        try:
            conn = sqlite3.connect(DB_PATH)
            cursor = conn.cursor()

            #qry = "SELECT sha, date_test, name_group, result, COUNT() FROM logs GROUP BY sha, name_group \
            #    ORDER BY date_test DESC LIMIT 10;".format()

            qry = "SELECT substr(sha, -8) as commit_sha, date_test, name_group, result, sum(case when result='PASS' then 1 else 0 end), COUNT(*) FROM logs\
                GROUP BY name_group, commit_sha ORDER BY date_test DESC;".format()

            cursor.execute(qry)
            result_qry = cursor.fetchall()

        except Exception as e:
            print(str(e))
            conn.rollback()
        finally:
            cursor.close()
            conn.close()
            return result_qry
