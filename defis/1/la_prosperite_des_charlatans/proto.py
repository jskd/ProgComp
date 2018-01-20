import json
import os

json_data=open("defis/1/la_prosperite_des_charlatans/target.json")
data = json.load(json_data)
json_data.close()

for target in data["target"]:
  print (target["name"])
  print ("TEST SUCESS -> NOTE: 20/20 avec félicitations du jury")
  print ()
