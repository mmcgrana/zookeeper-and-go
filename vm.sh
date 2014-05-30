#!/bin/bash

# Install Docker in the VM.
wget -q -O - https://get.docker.io/gpg | apt-key add -
echo deb http://get.docker.io/ubuntu docker main > /etc/apt/sources.list.d/docker.list
apt-get update
apt-get install -y --force-yes lxc-docker

# Build the Docker images we'll be using.
docker build -t zookeeper /vagrant/img-zookeeper
docker build -t go /vagrant/img-go

# Determine how Zookeeper will address itself in the cluster and how
# clients will address the cluster.
VM_IP=$1
ZOOKEEPER_SERVERS_ZK="$VM_IP:12888:13888,$VM_IP:22888:23888,$VM_IP:32888:33888"
ZOOKEEPER_SERVERS_GO="$VM_IP:12181,$VM_IP:22181,$VM_IP:32181"

# Create persistent Zookeeper data directories.
mkdir -p /var/zookeeper.{1,2,3}

# Configure Upstart scripts for Zookeepers. Run the servers in
# Docker, give each server the relevant cluster discovery
# information, and bind them to persistent data directories.
cat > /etc/init/zookeeper.1.conf <<EOF
exec docker run --env ZOOKEEPER_ID=1 --env ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_ZK --env ZOOKEEPER_CLIENT_PORT=12181 --net host --volume /var/zookeeper.1:/var/zookeeper zookeeper
EOF

cat > /etc/init/zookeeper.2.conf <<EOF
exec docker run --env ZOOKEEPER_ID=2 --env ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_ZK --env ZOOKEEPER_CLIENT_PORT=22181 --net host --volume /var/zookeeper.2:/var/zookeeper zookeeper
EOF

cat > /etc/init/zookeeper.3.conf <<EOF
exec docker run --env ZOOKEEPER_ID=2 --env ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_ZK --env ZOOKEEPER_CLIENT_PORT=32181 --net host --volume /var/zookeeper.2:/var/zookeeper zookeeper
EOF

# Start Zookeepers.
start zookeeper.1
start zookeeper.2
start zookeeper.3

# Write Zookeeper client addresses for use by Go clients.
cat > /etc/go.env <<EOF
ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_GO
EOF
