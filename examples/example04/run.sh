#!/bin/sh

export PWD=`pwd`
export CGO_CFLAGS="-I ${PWD}/include01/header"
export LD_LIBRARY_PATH="${PWD}/include01/source"
export CGO_LDFLAGS="-L ${PWD}/include01/source"

go run .