FROM golang

ADD . /go/src/test/test3

RUN go install test/test3/main

ENV GODEBUG=gctrace=1

ENTRYPOINT /go/bin/main
