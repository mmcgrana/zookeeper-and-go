#!/usr/bin/env bash

sudo apt-get update
sudo apt-get install -y openjdk-7-jdk
sudo apt-get install -y zookeeper
/usr/share/zookeeper/bin/zkServer.sh start
