#!/bin/bash

#用service将httpserver发布给集群内部访问
kubectl create -f httpserver-service.yaml

#安装helm v3版本
wget https://get.helm.sh/helm-v3.6.0-linux-amd64.tar.gz

#解压
tar -zxvf helm-v3.6.0-linux-amd64.tar.gz

# 移动到环境变量目录里面即可
mv linux-amd64/helm /usr/local/bin/helm

# 输出版本
helm version

#创建一个chart，chart的名称叫 helm-test
helm create helm-test
cd helm-test/

# 添加 ingress-nginx 官方helm仓库
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
cd /usr/local/src/

# # 查找ingress-nginx的chart包
helm search repo ingress-nginx

# 下载ingress-nginx
helm pull ingress-nginx/ingress-nginx

# 解压
tar -zxvf ingress-nginx-4.0.10.tgz 

# 进入ingress-nginx目录
cd /ingress-nginx

#下载阿里云镜像
docker pull registry.aliyuncs.com/google_containers/nginx-ingress-controller:v1.1.0
docker tag registry.aliyuncs.com/google_containers/nginx-ingress-controller:v1.1.0 k8s.gcr.io/ingress-nginx/controller:v1.1.0
docker rmi registry.aliyuncs.com/google_containers/nginx-ingress-controller:v1.1.1
docker pull registry.aliyuncs.com/google_containers/kube-webhook-certgen:v1.1.1
docker tag registry.aliyuncs.com/google_containers/kube-webhook-certgen:v1.1.1 k8s.gcr.io/ingress-nginx/kube-webhook-certgen:v1.1.1
docker rmi registry.aliyuncs.com/google_containers/kube-webhook-certgen:v1.1.1

#命名空间
kubectl create ns ingress-nginx

#helm安装ingress-nginx
cd /usr/local/src/ingress-nginx
helm install ingress-nginx -n ingress-nginx .

#通过创建tls.crt和tls.key来确保httpservice的安全
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=cncamp.com/O=cncamp"

#解析tls.crt和tls.key,并替换httpserver-secret.yaml里的tls.crt和tls.key
cat tls.crt|base64 -w 0
cat tls.key|base64 -w 0

#创建httpserver-ingress发布至到集群外部，供外部访问
kubectl apply -f httpserver-ingress.yaml
