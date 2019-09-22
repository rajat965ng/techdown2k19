#!/usr/bin/env bash
sudo adduser --disabled-password --gecos \packer\ packer;
su packer;
sudo apt-get update -y;
sudo apt-get install default-jdk -y;
wget http://mirrors.estointernet.in/apache/tomcat/tomcat-8/v8.5.46/bin/apache-tomcat-8.5.46.tar.gz;
tar -xvf apache-tomcat-8.5.46.tar.gz;
sudo apt-get install docker.io -y;
sudo gpasswd -a $(whoami) docker;
newgrp docker;
docker pull nginx;