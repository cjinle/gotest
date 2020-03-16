#!/bin/bash

# https://github.com/lonsharn/leafserver

export GOPATH=`pwd`
go build -o bin/server -gcflags "-N -l" src/server/main.go
go build -o bin/client -gcflags "-N -l" src/test/client.go
go build -o bin/test_glog -gcflags "-N -l" src/test/test_glog.go
