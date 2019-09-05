#!/bin/bash

kubectl apply -f namespace/micro-test.yaml
kubectl apply -f service/micro-test-service.yaml
kubectl apply -f deployment/micro-test-deployment.yaml
