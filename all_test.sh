#!/bin/bash

OPTION=$1

go clean -testcache
go test $OPTION ./test...
go test $OPTION ./db...
