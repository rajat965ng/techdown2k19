# K8 Jenkins Cluster

## Pre-Requisite
    Up and running 3 node Kubernetes cluster.
    Configured kubectl with cluster. Verify by executing 'kubectl get nodes'.
    Create docker image of Jenkins Master using 'on-prem/k8/jenkins/master/Dockerfile' and host it on a docker hub. Eg. rajat965ng/jenkins-master
    Create docker image of Jenkins Agent using 'on-prem/k8/jenkins/agent/Dockerfile' and host it on a docker hub. Eg. rajat965ng/jnlp-docker
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

## Configure Jenkins instance
    Go to 'Manage Jenkins -> Configure System' 
           # of executors: 0
    Go to 'Docker Slaves'
           Maximum number of running docker-slaves: 100
    Go to 'Cloud -> Add a new cloud -> Kubernetes'
    Go to 'Pod Labels -> Pod Label'
           key: jenkins
           value: slave
    Go to 'Pod Templates -> Add Pod Template'
            Name: jenkins-slave
            Label: jenkins-slave
            containers -> container template
                       Name: jnlp
                       docker image: rajat965ng/jnlp-docker
                       working directory: /home/jenkins
                       command to run: ''
                       arguments to pass the command: ''
                       EnvVars -> Environment Variable
                                  key: JENKINS_URL
                                  value: http://jenkins-master:8080/jenkins                    
            volumes -> host path volumes
                       host path: /var/run/docker.sock
                       mount path: /var/run/docker.sock    