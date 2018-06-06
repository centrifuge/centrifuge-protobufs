# centrifuge-documents
Protobuf files &amp; go bindings for Centrifuge Documents

[![Build Status](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs.svg?token=Sbf68xBZUZLMB3kGTKcX&branch=master)](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs)

# How to commit
Make sure you run the following two commands before committing to make sure the generated files are added.
```bash,
prototool all
npm run build_swagger
```

# Installation

## Making sure all dependencies are installed

```bash,
glide install
go install github.com/CentrifugeInc/centrifuge-protobufs/vendor/github.com/roboll/go-vendorinstall

PATH=$PATH:$GOPATH/bin

go-vendorinstall github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go-vendorinstall github.com/golang/protobuf/protoc-gen-go
go-vendorinstall github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

npm install
```

## Install Prototool
Install prototool with:

```
curl -sSL https://github.com/uber/prototool/releases/download/v0.2.0/prototool-$(uname -s)-$(uname -m) \
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

# To build the swagger html docs
npm run build_swagger

```


# Swagger
To view the swagger documentation, you can execute the following two commands:

```bash,
npm run build_swagger
open gen/swagger/html/index.html
```


