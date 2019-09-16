- [Metadata Script]
    - sudo mkdir -p /var/opt/jfrog/artifactory
    - sudo useradd artifactory
    - sudo usermod -aG docker artifactory
    - sudo chmod -R 777 /var/opt/jfrog/artifactory
    - docker run -d -v /var/opt/jfrog/artifactory:/var/opt/jfrog/artifactory -p 80:8081 docker.bintray.io/jfrog/artifactory-pro:latest

- [Jfrog Server: Http Settings (http://104.198.239.157)]
    - Docker Access Method: Repository Path
    - Server Provider: Embedded Tomcat

- [Jfrog artifactory image tag/push]
    - docker login -u admin -p password 104.198.239.157:80
    - docker tag 01176385d84a 104.198.239.157:80/docker-local/curl
    - docker push 104.198.239.157:80/docker-local/curl


- [For Setting Xray Create a seperate instance (http://34.68.148.114:8000)]
    - use file "xray"
    
    - ./xray install
    - ./xray start

- [To build the provider GO code]
    - go build -o terraform-provider-xray_v0.1.0


- [To Enable Logging]
    - export TF_LOG=DEBUG
    - export TF_TF_LOG_PATH=/Users/rajnigam/workspace/tf_jfrog/provisioner
    - terraform apply
