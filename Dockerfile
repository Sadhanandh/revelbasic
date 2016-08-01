FROM golang

COPY . /go/src/github.com/sadhanandh
WORKDIR /go/src/github.com/sadhanandh

# Install Revel CLI
RUN go get -u github.com/revel/cmd/revel

# Install project dependencies
RUN go get gopkg.in/mgo.v2

RUN go install ./revelbasic

EXPOSE 9000
ENTRYPOINT revel run github.com/sadhanandh/revelbasic/application devdocker 9000
