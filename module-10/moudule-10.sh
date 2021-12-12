#!/bin/bash

#下载prometheus-operator源码
git clone https://github.com/prometheus-operator/kube-prometheus.git 
cd kube-prometheus/
git checkout release-0.7
git branch

#添加CRD
cd kube-prometheus/manifests
kubectl apply -f setup/
kubectl apply -f .

#部署httpserver.yaml及svc
kubectl apply -f /opt/go/httpserver-deployment.yaml
kubectl apply -f /opt/go/httpserver-svc.yaml

#因prometheus-operator自定义指标需创建serivceMonitoring及prometheus-service
kubectl apply -f /kube-prometheus/manifests/prometheus-serviceMonitoringhttpserver.yaml
kubectl apply -f /kube-prometheus/manifests/prometheus-httpserver-service.yaml
