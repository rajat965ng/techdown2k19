# Gitea

## Pre-Requisite
    Up and running 3 node Kubernetes cluster.
    Configured kubectl with cluster. Verify by executing 'kubectl get nodes'.
    NFS for writing state.

## Configure 'deployment.yml'
### Configure nfs section in volumes
    volumes:
      - name: git-data
        nfs:
          path: /mnt/disk/gitea/data/app
          server: 10.150.16.171                

    volumes:
      - name: pgdata
        nfs:
          path: /mnt/disk/gitea/data/db
          server: 10.150.16.171                

## Installation
    cd on-prem/k8/git/
    kubectl apply -f deployment.yml

## Configure Giteas instance
### Setup admin username/password
    Username: abc
    Password: 8\Pd{H\Wt'w*R!Z=

## Configure Git Client
    git config --global http.sslVerify false
    git config --global user.email "abc@example.com"
    git config --global user.name "abc master"

