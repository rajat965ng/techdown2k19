#!/usr/bin/env bash

echo 'y' | sudo kubeadm reset
sudo yum remove -y kubeadm kubectl kubelet kubernetes-cni kube*
sudo yum -y autoremove
sudo rm -rf ~/.kube