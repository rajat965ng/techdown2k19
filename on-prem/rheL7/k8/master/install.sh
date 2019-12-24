#!/usr/bin/env bash

echo `hostname -I | awk '{ print $1 }'` >> ip.out
sudo kubeadm init  --apiserver-advertise-address=`cat ip.out` --pod-network-cidr=10.244.0.0/16

mkdir -p $HOME/.kube && sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config && sudo chown $(id -u):$(id -g) $HOME/.kube/config

#sudo yum install wget -y && wget https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl apply -f kube-flannel.yml && sudo kubeadm token create --print-join-command

echo "Enable Ingress on Bare metal"
#kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml
kubectl apply -f mandatory.yaml

echo "Verify ingress-nginx (NLB)"
#kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/baremetal/service-nodeport.yaml
kubectl apply -f service-nodeport.yaml

echo "Verify installation"
kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx --watch