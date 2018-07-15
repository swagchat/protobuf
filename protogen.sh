#!/bin/bash

rm -f ./*.gw.go
rm -f ./*.pb.go
rm -f ./*.gen.go

protoc \
  -I./ \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  -I$GOPATH/src/github.com/gogo/protobuf \
  --gogo_out=plugins=grpc:../../../ \
  --grpc-gateway_out=logtostderr=true:../../../ \
  chatMessage.proto \
  chatService.proto \
  deviceMessage.proto \
  commonMessage.proto \
  roomMessage.proto \
  roomUserMessage.proto \
  roomUserService.proto \
  userMessage.proto \
  userService.proto \
  userRoleMessage.proto \
  userRoleService.proto

