#!/usr/bin/env bash
sudo adduser --disabled-password --gecos "packer" packer;
su packer;
export DEBIAN_FRONTEND=noninteractive;
sudo apt-get update -y;
ln -fs /usr/share/zoneinfo/America/New_York /etc/localtime;
sudo apt-get install -y tzdata;
sudo dpkg-reconfigure --frontend noninteractive tzdata;