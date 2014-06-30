FROM google/golang

WORKDIR /gopath/src/github.com/coreos/updatectl
ADD . /gopath/src/github.com/coreos/updatectl
RUN go get github.com/coreos/updatectl

CMD []
ENTRYPOINT ["/gopath/bin/updatectl"]
