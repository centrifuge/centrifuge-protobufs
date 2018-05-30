# centrifuge-documents
Protobuf files &amp; go bindings for Centrifuge Documents

[![Build Status](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs.svg?token=Sbf68xBZUZLMB3kGTKcX&branch=master)](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs)


# Installation

## Making sure all dependencies are installed

```bash
glide install
go install github.com/CentrifugeInc/centrifuge-protobufs/vendor/github.com/roboll/go-vendorinstall

PATH=$PATH:$GOPATH/bin

go-vendorinstall github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go-vendorinstall github.com/golang/protobuf/protoc-gen-go
go-vendorinstall github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```


## Install Prototool
Install prototool with:

```
curl -sSL https://github.com/uber/prototool/releases/download/v0.1.0/prototool-$(uname -s)-$(uname -m) \
  -o /usr/local/bin/prototool && \
  chmod +x /usr/local/bin/prototool && \
  prototool -h
```

# Quick intro to prototool
We are using [prototool](https://github.com/uber/prototool) to lint or protobuf
files and build our go stubs.

## Helpful commands

Below a few helpful commands:

```
# To lint your files:
prototool lint

# To check if they compile:
prototool compile

# To build the go stubs:
prototool gen
```