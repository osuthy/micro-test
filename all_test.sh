#!/bin/bash

go clean -testcache
go test -v ./test
go test -v ./db
