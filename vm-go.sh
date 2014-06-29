#!/usr/bin/env bash

GO_VERSION="1.2.2"
GO_ARCHIVE="go${GO_VERSION}.linux-amd64.tar.gz"

apt-get update
apt-get install -y build-essential git-core mercurial
cd /opt
wget -q "https://storage.googleapis.com/golang/$GO_ARCHIVE"
tar -xzf "$GO_ARCHIVE"

cat > /home/vagrant/.profile <<EOF
export GOROOT=/opt/go
export GOPATH=\$HOME
export PATH=\$HOME/bin:\$GOROOT/bin:\$PATH
export ZOOKEEPER_SERVERS=192.168.12.11:2181,192.168.12.12:2181,192.168.12.13:2181
EOF
chown vagrant:vagrant /home/vagrant/.profile

sudo -u vagrant -i go get github.com/samuel/go-zookeeper/zk
sudo -u vagrant -i go get github.com/mmcgrana/zk
