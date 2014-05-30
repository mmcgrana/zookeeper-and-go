#!/bin/bash

if [ -z "$ZOOKEEPER_ID" ]; then
    echo "Need ZOOKEEPER_ID"
    exit 1
fi

if [ -z "$ZOOKEEPER_SERVERS" ]; then
    echo "Need ZOOKEEPER_SERVERS"
    exit 1
fi

if [ -z "$ZOOKEEPER_CLIENT_PORT" ]; then
    echo "Need ZOOKEEPER_CLIENT_PORT"
    exit 1
fi

mkdir -p /var/zookeeper/{data,conf}

echo -n $ZOOKEEPER_ID > /var/zookeeper/data/myid

cat > /var/zookeeper/conf/zoo.cfg <<EOF
tickTime=2000
initLimit=10
syncLimit=5
dataDir=/var/zookeeper/data
EOF

echo  "clientPort=$ZOOKEEPER_CLIENT_PORT" >> /var/zookeeper/conf/zoo.cfg

ZOOKEEPER_SERVERS=(${ZOOKEEPER_SERVERS//,/ })
for INDEX in "${!ZOOKEEPER_SERVERS[@]}"
do
    ID=$(expr $INDEX + 1)
    echo "server.$ID=${ZOOKEEPER_SERVERS[INDEX]}" >> /var/zookeeper/conf/zoo.cfg
done

exec /opt/zookeeper-3.4.6/bin/zkServer.sh start-foreground /var/zookeeper/conf/zoo.cfg
