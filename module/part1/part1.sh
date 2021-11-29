#!/bin/bash

#部署高可用httpserver，并设置探活及优雅终止
kubectl apply -f httpserver-deployment.yaml 
kubectl create -f no-gigterm.yaml  
