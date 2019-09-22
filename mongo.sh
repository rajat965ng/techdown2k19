#!/usr/bin/env bash
sudo apt-get install apt-transport-https -y
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add - ;
echo "deb [ arch=amd64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list ;
sudo apt-get update -y;
sudo apt-get install -y mongodb-org;
echo "mongodb-org hold" | sudo dpkg --set-selections ;
echo "mongodb-org-server hold" | sudo dpkg --set-selections ;
echo "mongodb-org-shell hold" | sudo dpkg --set-selections ;
echo "mongodb-org-mongos hold" | sudo dpkg --set-selections ;
echo "mongodb-org-tools hold" | sudo dpkg --set-selections ;