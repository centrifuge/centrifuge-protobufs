# centrifuge-documents
Protobuf files &amp; go bindings for Centrifuge Documents

[![Build Status](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs.svg?token=Sbf68xBZUZLMB3kGTKcX&branch=master)](https://travis-ci.com/CentrifugeInc/centrifuge-protobufs)

## Quick intro to prototool
We are using [prototool](https://github.com/uber/prototool) to lint or protobuf
files and build our go stubs. 

### Installation
Install prototool with:

```
curl -sSL https://github.com/uber/prototool/releases/download/v0.2.0/prototool-$(uname -s)-$(uname -m) \
  -o /usr/local/bin/prototool && \
  chmod +x /usr/local/bin/prototool && \
  prototool -h
```

### Helpful commands

Below a few helpful commands:

```
# To lint your files:
prototool lint

# To check if they compile:
prototool compile

# To build the go stubs:
prototool gen
```