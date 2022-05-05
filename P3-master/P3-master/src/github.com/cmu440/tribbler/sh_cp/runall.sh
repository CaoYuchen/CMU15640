#!/bin/bash

if [ -z $GOPATH ]; then
    echo "WARNING! GOPATH environment variable is not set!"
    exit 1
fi

if [ -n "$(go version | grep 'darwin/amd64')" ]; then    
    GOOS="darwin_amd64"
elif [ -n "$(go version | grep 'linux/amd64')" ]; then
    GOOS="linux_amd64"
elif [ -n "$(go version | grep 'darwin/arm64')" ]; then
    GOOS="darwin_amd64"
else
    echo "FAIL: only 64-bit Mac OS X and Linux operating systems are supported"
    exit 1
fi

$GOPATH/src/github.com/cmu440/tribbler/sh_cp/tribtest.sh
$GOPATH/src/github.com/cmu440/tribbler/sh_cp/libtest.sh
$GOPATH/src/github.com/cmu440/tribbler/sh_cp/storagetest.sh
