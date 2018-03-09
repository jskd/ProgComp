FROM openjdk:8u151 

ENV SCALA_VERSION 2.12.4 
ENV SBT_VERSION 1.1.1


RUN touch /usr/lib/jvm/java-8-openjdk-amd64/release

RUN \
  curl -fsL https://downloads.typesafe.com/scala/$SCALA_VERSION/scala-$SCALA_VERSION.tgz | tar xfz - -C /root/ && \
  echo >> /root/.bashrc && \
  echo "export PATH=~/scala-$SCALA_VERSION/bin:$PATH" >> /root/.bashrc


RUN \
  curl -L -o sbt-$SBT_VERSION.deb https://dl.bintray.com/sbt/debian/sbt-$SBT_VERSION.deb && \
  dpkg -i sbt-$SBT_VERSION.deb && \
  rm sbt-$SBT_VERSION.deb && \
  apt-get update && \
  apt-get install sbt && \
  sbt sbtVersion


WORKDIR /scableur
ADD . /scableur
RUN mkdir -p resources/input
RUN mkdir -p resources/output
RUN chmod +x run.sh
RUN sbt assembly