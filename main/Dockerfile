FROM golang:1.15.5-alpine AS build

WORKDIR /go/src/github.com/cjinle/test/main

ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

CMD ["./main"]
