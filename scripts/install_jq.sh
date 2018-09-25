#!/usr/bin/env bash

if [ -e ~/bin/jq ]; then
    echo "Found existing jq in bin. Not downloading again"
else
    os=`uname -s`
    echo "Downloading jq for $os"
    if [ "$os" = "Darwin" ]; then
        curl -sSL https://github.com/stedolan/jq/releases/download/jq-1.5/jq-osx-amd64 > ~/bin/jq
    else
        curl -sSL https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 > ~/bin/jq
    fi

    chmod +x ~/bin/jq
fi
