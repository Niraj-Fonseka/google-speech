#!/usr/bin/env bash

go build -i -v -o bin/server $GOPATH/src/google-speech/cmd/server

go build -i -v -o bin/client $GOPATH/src/google-speech/cmd/client
