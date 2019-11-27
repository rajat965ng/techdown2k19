#!/usr/bin/env bash

sudo yum install docker -y && sudo systemctl enable docker && sudo systemctl start docker

sudo setenforce 0 && sudo sed -i --follow-symlinks 's/^SELINUX=enforcing/SELINUX=disabled/' /etc/sysconfig/selinux


sudo systemctl disable firewalld && sudo systemctl stop firewalld

sudo sed -i '/swap/d' /etc/fstab
sudo swapoff -a


sudo tee -a /etc/sysctl.d/k8s.conf > /dev/null <<EOT
net.bridge.bridge-nf-call-ip6tables = 1
net.bridge.bridge-nf-call-iptables = 1
EOT

sudo sysctl --system

sudo tee -a /proc/sys/net/ipv4/ip_forward > /dev/null <<EOT
1
EOT

sudo tee -a /etc/yum.repos.d/kubernetes.repo > /dev/null <<EOT
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOT


sudo yum install -y kubelet kubeadm kubectl --disableexcludes=kubernetes && sudo systemctl enable --now kubelet && sudo systemctl daemon-reload && sudo systemctl restart kubelet