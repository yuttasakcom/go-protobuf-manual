# Go Protobuf Manual

## Install compiler

```bash
go get -u google.golang.org/protobuf
go get -u google.golang.org/protobuf/proto
```

## Install protocol compiler plugins for Go

> https://grpc.io/docs/languages/go/quickstart/#prerequisites

## Run Complier

```bash
$ protoc --go_out=. ./proto/basic/*.proto
```

## Proto Result

```proto
syntax = "proto3";

option go_package = "github.com/yuttasakcom/go-protobuf-manual";

message Hello {
  string name = 1;
}
```

## .zshrc

```bash
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOROOT:$GOPATH:$GOBIN
```
