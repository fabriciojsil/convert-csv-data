FROM golang:1.10.2-stretch

MAINTAINER fabriciojsil@gmail.com

ADD . $GOPATH/src/github.com/fabriciojsil/convert-csv-data
WORKDIR $GOPATH/src/github.com/fabriciojsil/convert-csv-data

RUN go get -u github.com/pilu/fresh

RUN make build

CMD ./csv-converter -file=$FILE -format=$FORMAT
