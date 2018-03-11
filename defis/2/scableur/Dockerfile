From p7hb/docker-spark

WORKDIR /scableur
ADD . /scableur

RUN mkdir -p resources/input
RUN mkdir -p resources/output

RUN chmod +x run.sh
RUN sbt package