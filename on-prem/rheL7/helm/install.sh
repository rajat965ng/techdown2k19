#!/usr/bin/env bash

sudo yum install curl -y;

sudo wget https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 --output-document=get_helm.sh
sudo chmod 700 get_helm.sh
sudo ./get_helm.sh

#helm repo add elasticsearch https://hub.helm.sh/charts/elastic/elasticsearch
#
#helm search repo elasticsearch https://hub.helm.sh/charts/elastic/elasticsearch
#
#
#helm install --name elasticsearch https://hub.helm.sh/charts/elastic/elasticsearch
#helm install filebeat https://hub.helm.sh/charts/elastic/filebeat
#helm install kibana https://hub.helm.sh/charts/elastic/kibana
#helm install metricbeat https://hub.helm.sh/charts/elastic/metricbeat