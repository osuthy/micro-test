#!/bin/bash

set -eu

cat deployment/micro-test-deployment.yaml | NAME_SPACE=$1 envsubst > deployment.yaml
kubectl apply -f deployment.yaml
