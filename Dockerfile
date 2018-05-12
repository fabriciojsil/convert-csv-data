FROM golang:1.10.2-stretch

MAINTAINER fabriciojsil@gmail.com

ADD . $GOPATH/src/github.com/fabriciojsil/convert-csv-data
WORKDIR $GOPATH/src/github.com/fabriciojsil/convert-csv-data

RUN go get -u github.com/pilu/fresh

RUN go build -o convert-csv-data -i cmd/server/main.go

# ENTRYPOINT ./convert-csv-data
