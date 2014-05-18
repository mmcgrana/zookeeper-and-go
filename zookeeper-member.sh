#!/usr/bin/env bash

apt-get update
apt-get install -y openjdk-7-jdk
wget -q -O /opt/zookeeper-3.4.6.tar.gz http://apache.mirrors.pair.com/zookeeper/zookeeper-3.4.6/zookeeper-3.4.6.tar.gz
tar -xzf /opt/zookeeper-3.4.6.tar.gz -C /opt

MYID=$1

mkdir -p /var/zookeeper/{data,conf}
echo -n $MYID > /var/zookeeper/data/myid
cat > /var/zookeeper/conf/zoo.cfg <<EOF
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/var/zookeeper/data
clientPort=2181
server.1=192.168.12.10:2888:3888
server.2=192.168.12.11:2888:3888
server.3=192.168.12.12:2888:3888
EOF
/opt/zookeeper-3.4.6/bin/zkServer.sh start /var/zookeeper/conf/zoo.cfg 
