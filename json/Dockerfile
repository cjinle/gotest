FROM scratch

WORKDIR $GOPATH/src/github.com/cjinle/test/json

COPY . $GOPATH/src/github.com/cjinle/test/json

CMD ["./main"]


# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .