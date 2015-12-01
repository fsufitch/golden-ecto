#!/bin/bash

export GOPATH=`dirname $BASH_SOURCE | xargs readlink -f`
echo "GOPATH=$GOPATH"

