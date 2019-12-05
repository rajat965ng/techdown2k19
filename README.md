# Platform

## Kubernetes Setup
### Pre-requisites
    Master VM
    Worker-A VM
    Worker-B VM
### Install Master 
    1. ssh into Master VM and switch to root user
    2. git clone -b kube/setup https://github.com/rajat965ng/techdown2k19.git
    3. cd techdown2k19
    4. sh -x on-prem/centOs7/k8/install.sh
    5. sh -x on-prem/centOs7/k8/master/install.sh
    6. execute command 'sudo kubeadm token create --print-join-command' and copy the output
### Install Worker-A/B VM 
    1. ssh into Worker VM and switch to root user
    2. git clone -b kube/setup https://github.com/rajat965ng/techdown2k19.git
    3. cd techdown2k19
    4. sh -x on-prem/centOs7/k8/install.sh    
    5. execute the output obtained from master

## Jenkins Setup
### Pre-requisites
    Jenkins VM
### Installation    
    1. ssh into Master VM and switch to root user
    2. git clone -b kube/setup https://github.com/rajat965ng/techdown2k19.git
    3. cd techdown2k19
    4. sh -x on-prem/centOs7/jenkins/install.sh
    5. copy $HOME/.kube/config from kubernetes Master VM and paste the content at $HOME/.kube/config in Jenkins VM


## Reverse Proxy Setup
### Pre-requisites
    Proxy VM
### Installation    
    1. ssh into Proxy VM and switch to root user
    2. git clone -b kube/setup https://github.com/rajat965ng/techdown2k19.git
    3. cd techdown2k19
    4. sh -x on-prem/centOs7/nginx/install.sh
    5. Follow the post-installation steps mentioned in on-prem/centOs7/nginx/README.md

## ELK Setup
### Pre-requisites
    A VM where kubectl is configured
### Installation    
    1. ssh into VM 
    2. git clone -b kube/setup https://github.com/rajat965ng/techdown2k19.git
    3. cd techdown2k19
    4. kubectl apply -f containers/elasticsearch/deployment.yml
    5. kubectl apply -f containers/logstash/deployment.yml
    6. kubectl apply -f containers/kibana/deployment.yml
