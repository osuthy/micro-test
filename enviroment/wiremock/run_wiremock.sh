#!/bin/bash

docker run -it --rm -v $(pwd):/home/wiremock -p 8080:8080 rodolpheche/wiremock
