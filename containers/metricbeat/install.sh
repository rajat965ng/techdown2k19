#!/usr/bin/env bash

sudo curl -L -O https://artifacts.elastic.co/downloads/beats/metricbeat/metricbeat-7.4.2-x86_64.rpm
sudo rpm -vi metricbeat-7.4.2-x86_64.rpm

sudo cp metricbeat.yml /etc/metricbeat/

sudo metricbeat modules enable kubernetes

sudo metricbeat setup
sudo service me tricbeat start