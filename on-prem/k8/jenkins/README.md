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

## Configure Kubectl CI plugin for CD [https://github.com/jenkinsci/kubernetes-cli-plugin]
    # Get the name of the token that was automatically generated for the ServiceAccount `jenkins-robot`.
    $ kubectl -n <namespace> get serviceaccount jenkins-robot -o go-template --template='{{range .secrets}}{{.name}}{{"\n"}}{{end}}'
    jenkins-robot-token-d6d8z
    
    # Retrieve the token and decode it using base64.
    $ kubectl -n <namespace> get secrets jenkins-robot-token-d6d8z -o go-template --template '{{index .data "token"}}' | base64 -d
    eyJhbGciOiJSUzI1NiIsImtpZCI6IiJ9.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2V[...]
    
    On Jenkins, navigate in the folder you want to add the token in, or go on the main page. 
    Then click on the "Credentials" item in the left menu and find or create the "Domain" you want. Finally, paste your token into a Secret text credential. 
    The ID is the credentialsId you need to use in the plugin configuration.

### Example:                             
    stage('List Pods'){
        steps {
            withKubeConfig(credentialsId: 'abc-dev',serverUrl: 'https://10.150.16.171:6443',clusterName: 'abccluster',contextName: 'abccluster') {
            sh 'kubectl get pods'
            }
        }
    }  
    stage('docker with registry') {
        steps{
           script {
               withDockerRegistry(credentialsId: 'localnexus', url: 'http://10.150.16.171:32000') {
                      docker.image('10.150.16.171:32000/docker-curl/node').inside {
                           sh 'node -v'
                      }
               }                        
           }
        }
    }
    stage('docker with registry but without script tag') {
        steps{
               withDockerRegistry(credentialsId: 'localnexus', url: 'http://10.150.16.171:32000') {
                      sh '''
                            docker pull 10.150.16.171:32000/docker-local/tutum/curl
                            docker tag  10.150.16.171:32000/docker-local/tutum/curl 10.150.16.171:32000/docker-curl/curl
                            docker push 10.150.16.171:32000/docker-curl/curl
                         '''
               }                        
        }
    }