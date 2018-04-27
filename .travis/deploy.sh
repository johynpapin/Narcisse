#!/usr/bin/env bash

go install $TRAVIS_BUILD_DIR
ssh johynpapin@78.248.188.78 'systemctl --user stop narcisse'
scp $GOPATH/bin/narcisse johynpapin@78.248.188.78:/home/johynpapin/narcisse
ssh johynpapin@78.248.188.78 'systemctl --user start narcisse'