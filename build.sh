#!/usr/bin/env bash

go build -v -o bin/server $GOPATH/src/google-speech/cmd/server

go build -v -o bin/client $GOPATH/src/google-speech/cmd/client
