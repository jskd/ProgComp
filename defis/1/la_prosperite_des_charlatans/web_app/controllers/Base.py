import inspect, sqlite3
from jinja2 import Environment, FileSystemLoader, TemplateNotFound
from site_config import site_config

DB_PATH = "data/database.db"

class BaseController:
    def get_template(self, caller_name):
        template = None
        try:
            class_name = self.__class__.__name__.replace('Controller', '')
            view_base = '{0}/views'.format(site_config.home)
            view_path = '{0}/{1}'.format(view_base, class_name.lower())
            env = Environment(loader=FileSystemLoader([view_base, view_path]))
            template = env.get_template(caller_name + '.html')
        except TemplateNotFound as template_ex:
            print(str(template_ex))
            pass
        return template

    def render_template(self, template_vars={}):
        try:
            caller_name = inspect.stack()[1][3]
            return self.get_template(caller_name).render(template_vars)
        except Exception as ex:
            return None

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
