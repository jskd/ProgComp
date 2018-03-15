#!/bin/bash

if [ "$1" == "--test" ]; then 

	if [ "$2" == "1" ]; then 
		echo "Running Test DataSet 1"

		DATA1="ressources/input1/data.csv"
		USERS1="ressources/input1/users.txt"
		VIEW1="output/view0.csv"
		CHANGES1="output/changes.txt"

		spark-submit target/scala-2.11/test_spark_2.11-0.1.jar $DATA1 $USERS1 $VIEW1 $CHANGES1
		exit 0;
	fi
	echo "No test was selected"
elif [ "$1" == "--clean" ]; then
	echo "Cleaning old output..."
	rm -rf output
else
	spark-submit target/scala-2.11/test_spark_2.11-0.1.jar $@
fi

