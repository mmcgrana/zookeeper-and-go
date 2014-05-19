#!/usr/bin/env bash

apt-get install -y python-software-properties
add-apt-repository -y ppa:duh/golang
apt-get update
apt-get install -y golang git-core

cat > /home/vagrant/.profile <<EOF
export GOPATH=\$HOME
export PATH=\$PATH:/home/vagrant/bin
export ZOOKEEPER_SERVERS=192.168.12.11:2181,192.168.12.12:2181,192.168.12.13:2181
EOF
chown vagrant:vagrant /home/vagrant/.profile
