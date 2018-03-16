From p7hb/docker-spark

WORKDIR /scableur
ADD . /scableur

RUN mkdir -p input
RUN touch input/data.csv
RUN touch input/users.txt

RUN mkdir -p output
RUN touch output/view0.csv
RUN touch output/changes.txt

RUN chmod +x run.sh
RUN sbt package