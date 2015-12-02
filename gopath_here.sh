#!/bin/bash

export GOPATH=`dirname $BASH_SOURCE | xargs readlink -f`
export GOBIN=$GOPATH/bin
echo "GOPATH=$GOPATH"
echo "GOBIN=$GOPATH/bin"

