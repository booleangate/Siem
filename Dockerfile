FROM golang

ADD . /go/src/github.com/stinkyfingers/Siem

RUN go get gopkg.in/mgo.v2
RUN go install github.com/stinkyfingers/Siem

ENTRYPOINT /go/bin/Siem

EXPOSE 8080