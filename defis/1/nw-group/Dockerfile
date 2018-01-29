FROM golang:stretch

WORKDIR /go/src
COPY Go/ Go/
RUN make -C Go
ENTRYPOINT ["Go/main.exe"]
