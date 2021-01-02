FROM golang:latest AS build

WORKDIR /go/src/github.com/cjinle/test/main

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

CMD ["./main"]
