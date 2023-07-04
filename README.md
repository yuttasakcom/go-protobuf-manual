# Go Protobuf Manual

## Install protocol compiler plugins for Go

> https://grpc.io/docs/languages/go/quickstart/#prerequisites

## Complier

```bash
$ protoc --go_out=. ./proto/basic/*.proto
```

## Hello Proto

```proto
syntax = "proto3";

option go_package = "github.com/yuttasakcom/go-protobuf-manual";

message Hello {
  string name = 1;
}
```
