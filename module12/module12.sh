#!/bin/sh

#下载安装istio
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.12.0
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y

#创建securesvc的namespace并打上label
kubectl create ns securesvc
kubectl label ns securesvc istio-injection=enabled

#创建httpserver.yaml
kubectl create -f httpserver.yaml -n securesvc

#为了https的安全性，创建key及cert证书，并创建istio-system的namespace并导入证书
openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -subj '/O=cncamp Inc./CN=*.cncamp.io' -keyout cncamp.io.key -out cncamp.io.crt

#创建istio-specs.yaml
kubectl apply -f istio-specs.yaml -n securesvc

