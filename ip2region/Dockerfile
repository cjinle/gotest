FROM scratch

WORKDIR /webapp

COPY ./ip2region.db ./main /webapp/

EXPOSE 8080

CMD ["./main"]

# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .



