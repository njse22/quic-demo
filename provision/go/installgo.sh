#!/bin/bash

curl -L -o go1.23.2.linux-amd64.tar.gz https://dl.google.com/go/go1.23.2.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.profile
sudo ldconfig
source .profile

sudo sysctl -w net.core.rmem_max=7500000
sudo sysctl -w net.core.wmem_max=7500000
