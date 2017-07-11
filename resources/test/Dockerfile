FROM golang:1.8-stretch

WORKDIR $GOPATH/src/github.com/gigmsn/subscriber/

ADD main.go .

ADD consumer/ consumer/

RUN go get -d -v ./...

ENTRYPOINT ["go", "test", "-v", "-race", "./..."]
