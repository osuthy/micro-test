#!/bin/bash

set -eu

TMP_FILE=$(mktemp)

cat deployment/micro-test-deployment.yaml | NAME_SPACE=$1 envsubst > $TMP_FILE
kubectl apply -f $TMP_FILE

cat $TMP_FILE
rm $TMP_FILE