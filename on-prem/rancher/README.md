# Rancher Kubernetes Engine (RKE)


## Pre-Requisites
### OS 
    RHEL 7/ CentOS 7
### Minimum number of nodes: 3
        IP          user    pass
    10.150.16.171   root    12345
    10.150.16.172   root    12345
    10.150.16.173   root    12345

### Install Docker on all 3 nodes 
    sh -x on-prem/rancher/docker/install.sh
    
### Create user and add password
    adduser gpssa
    passwd gpssa
    
        IP          user    pass
    10.150.16.171   gpssa   12345
    10.150.16.172   gpssa   12345
    10.150.16.173   gpssa   12345

### Make user to access docker without sudo
    groupadd docker
    chown gpssa:docker /var/run/docker.sock
    usermod -aG docker gpssa

### Enable password less login
    source machine: 10.150.16.171
    target machine: 10.150.16.171, 10.150.16.172, 10.150.16.173
    
    ACCESS SOURCE MACHINE AND PERFORM FOLLOWING STEPS
    
    Generate ssh token: 
        ssh-keygen -t rsa
    
    Push ssh tokens on target machine:
        ssh gpssa@10.150.16.171 mkdir -p .ssh
        cat .ssh/id_rsa.pub | ssh gpssa@10.150.16.171 'cat >> .ssh/authorized_keys'
        ssh gpssa@10.150.16.171 "chmod 700 .ssh; chmod 640 .ssh/authorized_keys"
        
        
        ssh gpssa@10.150.16.172 mkdir -p .ssh
        cat .ssh/id_rsa.pub | ssh gpssa@10.150.16.172 'cat >> .ssh/authorized_keys'
        ssh gpssa@10.150.16.172 "chmod 700 .ssh; chmod 640 .ssh/authorized_keys"
        
        
        ssh gpssa@10.150.16.173 mkdir -p .ssh
        cat .ssh/id_rsa.pub | ssh gpssa@10.150.16.173 'cat >> .ssh/authorized_keys'
        ssh gpssa@10.150.16.173 "chmod 700 .ssh; chmod 640 .ssh/authorized_keys"
        
## RKE Installation
    1. yum install wget -y
    2. wget https://github.com/rancher/rke/releases/download/v1.0.0/rke_linux-amd64
    3. mv rke_linux-amd64 rke
    4. chmod +x rke
    5. mv rke /usr/local/bin (move to $PATH) 
    6. rke --version (to verify the version)
    7. type 'ifconfig' and node down suitable ethernet interface eg. eth0, ens192 etc.
    8. update 'on-prem/rancher/k8/rancher-cluster.yml' with selected interface in network plugin section like 'canal_iface: ens192'
    9. execute 'rke up --config on-prem/rancher/k8/rancher-cluster.yml'
    10. export KUBECONFIG=$(pwd)/kube_config_rancher-cluster.yml
    11. yum install kubectl -y
    12. kubectl get nodes
    13. kubectl get pods --all-namespaces
    
## Notes
### To clean containers, images and resolve port issues
    docker kill $(docker ps -qa)
    docker rmi $(docker images -qa) --force
    docker rm  $(docker ps -qa)
    iptables -t filter -F
    iptables -t filter -X
    systemctl restart docker    

### Disable firewall
    yum install iptables-services
    systemctl start iptables
    systemctl enable iptables
    systemctl status iptables
    iptables -nvL
    systemctl stop iptables
    systemctl status iptables
            