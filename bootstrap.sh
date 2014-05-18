#!/usr/bin/env bash

apt-get update
apt-get install -y openjdk-7-jdk
wget -q -O /opt/zookeeper-3.4.6.tar.gz http://apache.mirrors.pair.com/zookeeper/zookeeper-3.4.6/zookeeper-3.4.6.tar.gz
tar -xzf /opt/zookeeper-3.4.6.tar.gz -C /opt

cat > /opt/zookeeper-3.4.6/conf/zoo.cfg <<EOF
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/tmp/zookeeper
clientPort=2181
EOF
/opt/zookeeper-3.4.6/bin/zkServer.sh start
