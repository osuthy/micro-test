#!/bin/bash

set -eu

cat namespace/micro-test-namespace.yaml | NAME_SPACE=$1 envsubst > namespace.yaml
kubectl apply -f namespace.yaml
