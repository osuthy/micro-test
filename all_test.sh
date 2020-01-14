#!/bin/bash

OPTION=$1

go clean -testcache
go test $OPTION ./testable/... | grep -v 'no test files'
go test $OPTION ./db/... | grep -v 'no test files'
go test $OPTION ./http/... | grep -v 'no test files'
go test $OPTION ./json/... | grep -v 'no test files'
go test $OPTION ./test/... | grep -v 'no test files'
