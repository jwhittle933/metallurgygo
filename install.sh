#!bin/bash

USE=$1
current=$(pwd)

function install_os() {
  git clone https://github.com/jwhittle933/metallurgygo \
    && cd ./metallurgygo/bin/
    && mv mgo /usr/local/bin/
    && cd $current \
    && rm -rf ./metallurgygo
}

function install_go() {
  go get github.com/jwhittle933/metallurgygo \
    && cd $GOPATH/src/github.com/jwhittle933/metallurgygo \
    && go install \
    && cd $GOPATH/bin/ \
    && mv metallurgygo mgo
}

if [ $USE ~= "os" ]; then
  install_os
else
  install_go
fi

