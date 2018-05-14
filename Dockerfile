FROM golang:1.10.2-stretch

MAINTAINER fabriciojsil@gmail.com

ADD . $GOPATH/src/github.com/fabriciojsil/convert-csv-data
WORKDIR $GOPATH/src/github.com/fabriciojsil/convert-csv-data

RUN make build

ENTRYPOINT ["docker-entrypoint.sh"]
