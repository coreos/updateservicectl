FROM google/golang

WORKDIR /gopath/src/github.com/coreos/updateservicectl
ADD . /gopath/src/github.com/coreos/updateservicectl
RUN go get github.com/coreos/updateservicectl

CMD []
ENTRYPOINT ["/gopath/bin/updateservicectl"]
