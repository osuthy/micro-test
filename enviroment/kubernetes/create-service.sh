#!/bin/bash

set -eu

cat service/micro-test-service.yaml | NAME_SPACE=$1 envsubst > service.yaml
kubectl apply -f service.yaml
