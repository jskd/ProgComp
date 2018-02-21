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

    def get_group_list(self):
        qry = "SELECT DISTINCT name_group FROM logs;"

        group_list = self.executeQuery(qry)
        return group_list

    def get_last_build_score(self, group_name):
        qry = "SELECT \
            sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
            COUNT(DISTINCT name_test) as total_test \
            FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
            WHERE name_group = '{}' \
            GROUP BY sha, name_group \
            ORDER BY date_test DESC LIMIT 1".format(group_name)

        last_build_score = self.executeQuery(qry)
        return last_build_score

    def get_last_valid_build_sha(self, group_name):
        qry = "SELECT sha FROM \
                (SELECT sha, \
                sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
                COUNT(DISTINCT name_test) as total_test \
                FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
                WHERE name_group = '{}' \
                GROUP BY sha, name_group \
                ORDER BY date_test DESC) \
            WHERE nb_passed = total_test \
            LIMIT 1;".format(group_name)

        last_valid_build_sha = self.executeQuery(qry)
        return last_valid_build_sha

    def get_group_stats(self):
        group_stats = self.get_group_list()

        for i, group in enumerate(group_stats):
            group_name = group[0]

            last_build_score = self.get_last_build_score(group_name)
            group_stats[i] += last_build_score[0]

            last_valid_build_sha = self.get_last_valid_build_sha(group_name)
            if last_valid_build_sha != [] :
                group_stats[i] += last_valid_build_sha[0]
            else: group_stats[i] += ('',)

        for row in group_stats:
            print(row)

        group_stats = sorted(group_stats, key=lambda x: x[1], reverse=True)
        return group_stats

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
