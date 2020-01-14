#!/bin/bash

set eu

eval $(minikube docker-env)
docker build -t micro-test-mysql .
