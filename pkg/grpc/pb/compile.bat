:: Install proto3.
:: https://github.com/google/protobuf/releases
:: Update protoc Go bindings via
::  go get -u github.com/golang/protobuf/proto
::  go get -u github.com/golang/protobuf/protoc-gen-go
::
:: See also
::  https://github.com/grpc/grpc-go/tree/master/examples

protoc location.proto --go_out=plugins=grpc:.