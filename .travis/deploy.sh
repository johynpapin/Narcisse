#!/usr/bin/env bash

go build $TRAVIS_BUILD_DIR
ssh johynpapin@78.248.188.78 'systemctl --user stop narcisse'
scp $TRAVIS_BUILD_DIR/narcisse johynpapin@78.248.188.78:/home/johynpapin/narcisse
ssh johynpapin@78.248.188.78 'systemctl --user start narcisse'