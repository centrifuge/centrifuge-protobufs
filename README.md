# centrifuge-documents
Protobuf files &amp; go bindings for Centrifuge Documents

[![Build Status](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs.svg?token=Sbf68xBZUZLMB3kGTKcX&branch=master)](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs)


# How to commit
Make sure you run the `make` before committing to make sure the generated files are up to date:

```bash,
# To install dependencies and generate files
make all
# To just generate files
make gen_proto gen_swagger
```

# Installation

```
$ make help
usage: make [target] ...

targets:
help                 Show this help message.
install              Install dependencies required to generate bindings & documentation
vendorinstall        Installs all protobuf dependencies with go-vendorinstall
lint                 runs prototool lint
gen_go               generates the go bindings
gen_proto            runs prototool all
gen_swagger          generates the swagger documenation
```
## Making sure all dependencies are installed
### Install Prototool
Install prototool with:

```
curl -sSL https://github.com/uber/prototool/releases/download/v0.4.0/prototool-$(uname -s)-$(uname -m) \
  -o /usr/local/bin/prototool && \
  chmod +x /usr/local/bin/prototool && \
  prototool -h
```

### Install dependencies
```
make install
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
# or: make gen_swagger
open gen/swagger/html/index.html
```


