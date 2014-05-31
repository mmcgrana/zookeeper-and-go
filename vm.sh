#!/bin/bash

# Install Docker in the VM.
wget -q -O - https://get.docker.io/gpg | apt-key add -
echo deb http://get.docker.io/ubuntu docker main > /etc/apt/sources.list.d/docker.list
apt-get update
apt-get install -y lxc-docker

# Install pipework.
apt-get install -y bridge-utils arping
wget -q -O /bin/pipework https://raw.githubusercontent.com/jpetazzo/pipework/master/pipework
chmod +x /bin/pipework

# Build the Docker images we'll be using.
docker build -t zookeeper /vagrant/img-zookeeper
docker build -t go /vagrant/img-go

# Determine how Zookeeper will address itself in the cluster and how
# clients will address the cluster.
ZOOKEEPER_1_IP="192.168.1.1"
ZOOKEEPER_2_IP="192.168.1.2"
ZOOKEEPER_3_IP="192.168.1.3"
GO_IP="192.168.1.4"
ZOOKEEPER_SERVERS_ZK="$ZOOKEEPER_1_IP:2888:3888,$ZOOKEEPER_2_IP:2888:3888,$ZOOKEEPER_3_IP:2888:3888"
ZOOKEEPER_SERVERS_GO="$ZOOKEEPER_1_IP:2181,$ZOOKEEPER_2_IP:2181,$ZOOKEEPER_3_IP:2181"

# Create persistent Zookeeper data directories.
mkdir -p /var/zookeeper.{1,2,3}

# Write environment data.
cat > /etc/zookeeper.1.env <<EOF
ZOOKEEPER_ID=1
ZOOKEEPER_CLIENT_PORT=2181
ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_ZK
EOF

cat > /etc/zookeeper.2.env <<EOF
ZOOKEEPER_ID=2
ZOOKEEPER_CLIENT_PORT=2181
ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_ZK
EOF

cat > /etc/zookeeper.3.env <<EOF
ZOOKEEPER_ID=3
ZOOKEEPER_CLIENT_PORT=2181
ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_ZK
EOF

cat > /etc/go.env <<EOF
ZOOKEEPER_SERVERS=$ZOOKEEPER_SERVERS_GO
EOF

# Start Zookeepers.
docker run -d --name zookeeper.1 --env-file /etc/zookeeper.1.env --volume /var/zookeeper.1:/var/zookeeper zookeeper
pipework br1 zookeeper.1 "$ZOOKEEPER_1_IP/24"

docker run -d --name zookeeper.2 --env-file /etc/zookeeper.2.env --volume /var/zookeeper.2:/var/zookeeper zookeeper
pipework br1 zookeeper.2 "$ZOOKEEPER_2_IP/24"

docker run -d --name zookeeper.3 --env-file /etc/zookeeper.3.env --volume /var/zookeeper.3:/var/zookeeper zookeeper
pipework br1 zookeeper.3 "$ZOOKEEPER_3_IP/24"

# Write go console script.
cat > /bin/go-console <<EOF
CONTAINER_ID=\$(docker run -d -t -i -v /vagrant:/vagrant --env-file /etc/go.env go /bin/bash)
pipework br1 \$CONTAINER_ID "$GO_IP/24"
docker attach \$CONTAINER_ID
EOF
chmod +x /bin/go-console
