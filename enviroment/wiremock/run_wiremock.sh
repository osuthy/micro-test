#!/bin/bash

cd `dirname $0`
docker run -it --rm -p 8080:8080 rodolpheche/wiremock
