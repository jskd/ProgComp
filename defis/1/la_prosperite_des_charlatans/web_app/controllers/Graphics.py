import cherrypy, sqlite3
from controllers.Base import BaseController
import os
import datetime
from string import Template

PATH_GRAPH_DATASET = "./web_app/content/graph_dataset/"
EXT_GRAPH_DATASET = ".graph-dataset"

PATH_GRAPH_CONFIG = "./web_app/content/graph_config/"
EXT_GRAPH_CONFIG = ".gnuplot"

PATH_GRAPH_RENDER = "./web_app/content/graph_render/"
EXT_GRAPH_RENDER = ".png"

GRAPH_CONFIG_TEPLATE = "./web_app/content/graph_config.template"

def get_dataset_path( idx ):
  return PATH_GRAPH_DATASET + str(idx) + EXT_GRAPH_DATASET

def get_label_dataset( name_test, teams ):
  label = '#Test: {}'.format(name_test) + os.linesep
  label = label + '{:19} {:40}'.format("#date", "sha_of_commit")
  for team in teams:
    label = label + ' {:10}'.format(team)
  return label

def get_line_dataset( date, sha, test_times ):
  line = '{:19} {:40}'.format(date, sha)
  for test_time in test_times:
    line = line + ' {:10}'.format(test_time)
  return line

def generate_graph_config( file_number, title, teams):
  template = open( GRAPH_CONFIG_TEPLATE )
  src = Template( template.read() )

  renderfile = PATH_GRAPH_RENDER + str(file_number) + EXT_GRAPH_RENDER
  datasetfile = PATH_GRAPH_DATASET + str(file_number) + EXT_GRAPH_DATASET

  plotconf = ""

  for idx, team in enumerate(teams):
    if idx != 0:
      plotconf += '  '
    plotconf += "'{}' using 1:{} title '{}' with linespoints".format(
      datasetfile, str(4+idx),team
    )
    if idx != len(teams)-1 :
      plotconf += ', \\'+ os.linesep

  result = src.substitute( {'renderfile' : renderfile, 'plotconf' : plotconf} )

  with open( PATH_GRAPH_CONFIG + str(file_number) + EXT_GRAPH_CONFIG ,'w') as f:
    f.write( result)

class GraphicsController(BaseController):
  @cherrypy.expose
  def index(self):

    list_commit_sha = self.all_sha_commit()
    self.generate_data_set()
    self.generate_config()

    teams = self.all_name_group()

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
      return '{:010.6f}'.format(self.executeQuery(qry)[0][0])
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
      return date_of_sha_commit
      #date_of_sha_commit = datetime.datetime.strptime(date_of_sha_commit, "%Y-%m-%d %H:%M:%S")
      #return str(date_of_sha_commit.strftime("%d/%m/%Y"))
    except:
      return "NaN"

  def generate_data_set(self):

    list_commit_sha = self.all_sha_commit()
    name_tests = self.all_name_test()
    groups = self.all_name_group()

    dataset_ext = ".txt"

    for idx, name_test in enumerate(name_tests):
      with open( get_dataset_path(idx) ,'w') as f:
        f.write( get_label_dataset(name_test, groups) + os.linesep)
        for sha in list_commit_sha:
          times = []
          for group in groups:
            time_result = self.get_time(name_test, sha, group)
            times.append(time_result)
          sha_date= self.get_date_of_sha_commit(sha)
          f.write(get_line_dataset(sha_date, sha, times) + os.linesep)


  def generate_config(self):
    name_tests = self.all_name_test()
    teams = self.all_name_group()

    for idx, name_test in enumerate(name_tests):
      generate_graph_config(idx, "name_test", teams)
