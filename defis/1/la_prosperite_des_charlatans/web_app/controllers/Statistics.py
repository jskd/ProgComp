import cherrypy, sqlite3
from controllers.Base import BaseController
import os

class StatisticsController(BaseController):
    @cherrypy.expose
    def index(self):
        lasts_commit_best_group = self.get_lasts_commit_best_group()
        group_stats = self.get_group_stats()

        template_args = {
        "last_commit_best_group":lasts_commit_best_group,
        "group_stats":group_stats
        }
        return self.render_template(template_args)

    @cherrypy.expose
    def history(self, group):
        lasts_commit_best_group = self.get_lasts_commit_best_group()
        group_history = self.get_history(group)

        template_args = {
        "last_commit_best_group":lasts_commit_best_group,
        "group_history":group_history
        }
        return self.render_template(template_args)


    def get_group_stats(self):

        qry = "SELECT sha as commit_sha, \
            date_test, \
            name_group, \
            sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
            COUNT(DISTINCT name_test) as total_test \
            FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
            WHERE commit_sha = (SELECT sha FROM logs ORDER BY date_test DESC LIMIT 1)\
            GROUP BY commit_sha, name_group \
            ORDER BY nb_passed DESC;".format()

        stats_by_group = self.executeQuery(qry)

        if stats_by_group != None:
            for i, row in enumerate(stats_by_group):
                group_name = row[2]

                qry = "SELECT * FROM \
                        (SELECT sha, \
                        date_test, \
                        name_group, \
                        sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
                        COUNT(DISTINCT name_test) as total_test \
                        FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
                        WHERE name_group = '{}' \
                        GROUP BY sha, name_group \
                        ORDER BY date_test DESC) \
                    WHERE nb_passed = total_test \
                    LIMIT 1;".format(group_name)

                last_valid_build = self.executeQuery(qry)

                valid_sha = ""

                if last_valid_build != None and len(last_valid_build) > 0:
                    valid_sha = last_valid_build[0][0]

                row += (valid_sha,)

                stats_by_group[i] = row

            for row in stats_by_group:
                print(row)

        return stats_by_group


    def get_history(self, group):

        qry = "SELECT sha, \
            date_test, \
            name_group, \
            sum(case when result='PASS' then 1 else 0 end) as nb_passed, \
            COUNT(DISTINCT name_test) as total_test \
            FROM (SELECT * FROM logs GROUP BY name_group, sha, name_test) \
            WHERE name_group = '{}' \
            GROUP BY sha, name_group \
            ORDER BY date_test DESC".format(group)

        return self.executeQuery(qry)
