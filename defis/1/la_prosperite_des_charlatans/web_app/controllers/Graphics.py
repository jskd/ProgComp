import cherrypy, sqlite3
from controllers.Base import BaseController
import os
import datetime

def get_label_dataset( name_test, teams ):
  label = '#Test: {}\n'.format(name_test)
  label = label + '{:10} {:40}'.format("#date", "sha_of_commit")
  for team in teams:
    label = label + ' {:10}'.format(team)
  return label

def get_line_dataset( date, sha, test_times ):
  line = '{:10} {:40}'.format(date, sha)
  for test_time in test_times:
    line = line + ' {:10}'.format(test_time)
  return line

class GraphicsController(BaseController):
  @cherrypy.expose
  def index(self):

    list_commit_sha = self.all_sha_commit()
    self.generate_data_set()

    template_args = {
    "data_set":list_commit_sha
    }
    return self.render_template(template_args)

  def all_sha_commit(self):
    qry = "SELECT distinct sha \
      FROM logs \
      ORDER BY date_test;".format()
    shas = []
    for (sha,) in self.executeQuery(qry):
      shas.append(sha)
    return shas

  def all_name_test(self):
    qry = "SELECT distinct name_test \
      FROM logs \
      ORDER BY name_test;".format()
    tests = []
    for (test,) in self.executeQuery(qry):
      tests.append(test)
    return tests

  def all_name_group(self):
    qry = "SELECT distinct name_group \
      FROM logs \
      ORDER BY name_group;".format()
    groups = []
    for (group,) in self.executeQuery(qry):
      groups.append(group)
    return groups

  def get_time(self, name_test, sha, name_group):
    qry = "SELECT bench_time \
      FROM logs \
      WHERE name_test = '{}' AND sha = '{}' AND name_group = '{}' AND result='PASS'\
      ORDER BY name_group \
      LIMIT 1;".format(
        name_test.replace("'","''"), sha, name_group)
    try:
      return '{:010.3f}'.format(self.executeQuery(qry)[0][0])
    except:
      return '{:10}'.format("NaN")

  def get_date_of_sha_commit(self, sha):
    qry = "SELECT date_test \
      FROM logs \
      WHERE sha = '{}' \
      ORDER BY date_test \
      LIMIT 1;".format(sha)
    try:
      date_of_sha_commit = self.executeQuery(qry)[0][0]
      date_of_sha_commit = datetime.datetime.strptime(date_of_sha_commit, "%Y-%m-%d %H:%M:%S")
      return str(date_of_sha_commit.strftime("%d/%m/%Y"))
    except:
      return "NaN"

  def generate_data_set(self):

    list_commit_sha = self.all_sha_commit()
    name_tests = self.all_name_test()
    groups = self.all_name_group()

    for name_test in name_tests:
      print( get_label_dataset(name_test, groups) )
      for sha in list_commit_sha:
        times = []
        for group in groups:
          time_result = self.get_time(name_test, sha, group)
          times.append(time_result)
        sha_date= self.get_date_of_sha_commit(sha)
        print(get_line_dataset(sha_date, sha, times))
