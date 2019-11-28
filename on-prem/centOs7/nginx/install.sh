#!/usr/bin/env bash

sudo yum install epel-release -y
sudo yum install nginx -y
sudo systemctl enable nginx
sudo systemctl start nginx
sudo systemctl status nginx
sudo systemctl reload nginx

sudo firewall-cmd --permanent --zone=public --add-service=http
sudo firewall-cmd --permanent --zone=public --add-service=https
sudo firewall-cmd --reload

