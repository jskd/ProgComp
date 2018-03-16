IMG="scableur-spark-img"

all: image

package:
	sbt package

clean:
	sbt clean

image:
	docker build -t $(IMG) .

mrproper:
	rm -rf target
	rm -rf project/target
	rm -rf project/project
	rm -rf *.class