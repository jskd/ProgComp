FROM golang:stretch

WORKDIR /go/src
COPY . .
RUN apt-get update && apt-get -y install ocaml
RUN make -C Go
ENTRYPOINT ["Go/main.exe"]
