#!/bin/sh

echo "bash arguments : $@"
if [ -f "out/Scableur.jar" ] ; then
	java -jar out/Scableur.jar $@
else
	echo "Sources not found.\nRun make command to compile program."
fi