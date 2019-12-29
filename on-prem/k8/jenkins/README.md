# K8 Jenkins Cluster

## Pre-Requisite
    Up and running 3 node Kubernetes cluster.
    Configured kubectl with cluster. Verify by executing 'kubectl get nodes'.
    Create docker image of Jenkins Master using 'on-prem/k8/jenkins/master/Dockerfile' and host it on a docker hub.
    Create docker image of Jenkins Agent using 'on-prem/k8/jenkins/agent/Dockerfile' and host it on a docker hub.
    NFS for writing state.

## Configure 'deployment.yml'
### Configure nfs section in volumes
     volumes:
       - name: docker-socket
         hostPath:
           path: /var/run/docker.sock
       - name: pvc
         nfs:
           path: /mnt/disk/vol
           server: 10.150.16.171     
### Configure jenkins image name in container section
     containers:
        - name:  jenkins
          image: rajat965ng/jenkins-master
    
## Installation
    cd on-prem/k8/jenkins/
    kubectl apply -f deployment.yml
