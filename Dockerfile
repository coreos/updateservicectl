FROM golang:1.6

WORKDIR /gopath/src/github.com/coreos/updateservicectl
ADD . /gopath/src/github.com/coreos/updateservicectl
RUN go get github.com/coreos/updateservicectl

CMD []
ENTRYPOINT ["/go/bin/updateservicectl"]
