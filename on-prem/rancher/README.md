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
    adduser abc
    passwd abc
    
        IP          user    pass
    10.150.16.171   abc   12345
    10.150.16.172   abc   12345
    10.150.16.173   abc   12345

### Make user to access docker without sudo
    groupadd docker
    chown abc:docker /var/run/docker.sock
    usermod -aG docker abc

### Enable password less login
    source machine: 10.150.16.171
    target machine: 10.150.16.171, 10.150.16.172, 10.150.16.173
    
    ACCESS SOURCE MACHINE AND PERFORM FOLLOWING STEPS
    
    Generate ssh token: 
        ssh-keygen -t rsa
    
    Push ssh tokens on target machine:
        ssh abc@10.150.16.171 mkdir -p .ssh
        cat .ssh/id_rsa.pub | ssh abc@10.150.16.171 'cat >> .ssh/authorized_keys'
        ssh abc@10.150.16.171 "chmod 700 .ssh; chmod 640 .ssh/authorized_keys"
        
        
        ssh abc@10.150.16.172 mkdir -p .ssh
        cat .ssh/id_rsa.pub | ssh abc@10.150.16.172 'cat >> .ssh/authorized_keys'
        ssh abc@10.150.16.172 "chmod 700 .ssh; chmod 640 .ssh/authorized_keys"
        
        
        ssh abc@10.150.16.173 mkdir -p .ssh
        cat .ssh/id_rsa.pub | ssh abc@10.150.16.173 'cat >> .ssh/authorized_keys'
        ssh abc@10.150.16.173 "chmod 700 .ssh; chmod 640 .ssh/authorized_keys"
        
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

## Setup SSL in RKE NGINX
    openssl genrsa 2048 > host.key
    chmod 400 host.key
    openssl req -new -x509 -nodes -sha256 -days 365 -key host.key -out host.cert
    kubectl -n ingress-nginx create secret tls ingress-default-cert --cert=host.cert --key=host.key -o yaml --dry-run=true > ingress-default-cert.yaml
    
    Include the contents of ingress-default-cert.yml inline with your RKE cluster.yml file. For example:
    
        addons: |-
          ---
          apiVersion: v1
          data:
            tls.crt: [ENCODED CERT]
            tls.key: [ENCODED KEY]
          kind: Secret
          metadata:
            creationTimestamp: null
            name: ingress-default-cert
            namespace: ingress-nginx
          type: kubernetes.io/tls
      
    Define your ingress resource with the following default-ssl-certificate argument, which references the secret we created earlier under extra_args in your cluster.yml:
    
        ingress: 
          provider: "nginx"
          extra_args:
            default-ssl-certificate: "ingress-nginx/ingress-default-cert"
        
    kubectl delete pod -l app=ingress-nginx -n ingress-nginx    
    rke up      

## Setup SSL in Rancher Server 
               
    docker run -v $PWD/certs:/certs \
               -e CA_SUBJECT="My own root CA" \
               -e CA_EXPIRE="1825" \
               -e SSL_EXPIRE="365" \
               -e SSL_SUBJECT="abc.example.com" \
               -e SSL_DNS="abc.example.com" \
               -e SILENT="true" \
               superseb/omgwtfssl           
               
               
    docker run -d --restart=unless-stopped \
               -p 80:80 -p 443:443 \
               -v $PWD/rancher:/var/lib/rancher \
               -v $PWD/certs/cert.pem:/etc/rancher/ssl/cert.pem \
               -v $PWD/certs/key.pem:/etc/rancher/ssl/key.pem \
               -v $PWD/certs/ca.pem:/etc/rancher/ssl/cacerts.pem \
               rancher/rancher:latest            
               
    docker run --rm --net=host superseb/rancher-check \
                                 "https://abc.example.com"
    
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
            
    
    
## Full length certs            
    openssl genrsa -des3 -out rootCA.key 2048      
    openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 1024  -out rootCA.pem
        
        Country Name (2 letter code) [AU]:NL
        State or Province Name (full name) [Some-State]:.
        Locality Name (eg, city) []:ACME City
        Organization Name (eg, company) [Internet Widgits Pty Ltd]:ACME Websites
        Organizational Unit Name (eg, section) []:ACME IT
        Common Name (e.g. server FQDN or YOUR name) []:ACME ROOT CA
        Email Address []:webmaster@acme.dev
    
    vi v3.ext
    
        authorityKeyIdentifier=keyid,issuer
        basicConstraints=CA:FALSE
        keyUsage = digitalSignature, nonRepudiation, keyEncipherment, dataEncipherment
        subjectAltName = @alt_names
        [alt_names]
        DNS.1 = acme-site.dev
        DNS.2 = acme-static.dev
        
    openssl req -new -nodes -out server.csr -newkey rsa:2048 -keyout server.key
    
        Country Name (2 letter code) [AU]:NL
        State or Province Name (full name) [Some-State]:.
        Locality Name (eg, city) []:ACME City
        Organization Name (eg, company) [Internet Widgits Pty Ltd]:ACME Websites
        Organizational Unit Name (eg, section) []:ACME IT
        Common Name (e.g. server FQDN or YOUR name) []:ACME DEV CERTIFICATE
        Email Address []:webmaster@acme.dev
        Please enter the following 'extra' attributes
        to be sent with your certificate request
        A challenge password []:ACME
        An optional company name []:.
        
    openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500 -sha256 -extfile v3.ext
    
            
    kubectl -n ingress-nginx create secret tls ingress-default-cert --cert=server.crt --key=server.key -o yaml --dry-run=true > ingress-default-cert.yaml
      