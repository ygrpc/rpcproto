#!/bin/bash

protoc -I=. --go_out=. --go-vtproto_out=. --go-vtproto_opt=features=marshal+unmarshal+size+pool proto.proto

cp -f github.com/ygrpc/rpcproto/*.pb.go .

rm -rf github.com