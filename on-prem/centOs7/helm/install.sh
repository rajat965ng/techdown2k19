#!/usr/bin/env bash

sudo yum install curl -y;

sudo wget https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 --output-document=get_helm.sh
sudo chmod 700 get_helm.sh
sudo ./get_helm.sh