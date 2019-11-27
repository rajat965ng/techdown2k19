#!/usr/bin/env bash

sudo yum install java-1.8.0-openjdk-devel curl -y
sudo yum install docker -y && sudo systemctl enable docker && sudo systemctl start docker

curl --silent --location http://pkg.jenkins-ci.org/redhat-stable/jenkins.repo | sudo tee /etc/yum.repos.d/jenkins.repo

sudo rpm --import https://jenkins-ci.org/redhat/jenkins-ci.org.key

sudo yum install jenkins -y
sudo systemctl start jenkins
sudo systemctl status jenkins
sudo systemctl enable jenkins

sudo groupadd docker
sudo chown jenkins:docker /var/run/docker.sock
sudo usermod -aG docker jenkins