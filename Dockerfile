FROM google/golang

WORKDIR /gopath/src/github.com/coreos-inc/updatectl
ADD . /gopath/src/github.com/coreos-inc/updatectl
RUN go get github.com/coreos-inc/updatectl

CMD []
ENTRYPOINT ["/gopath/bin/updatectl"]
