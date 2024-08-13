#!/bin/bash

protoc -I=. --go_out=. proto.proto

cp -f github.com/ygrpc/rpcproto/*.pb.go .

rm -rf github.com