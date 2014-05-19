## Zookeeper and Go

Source code and Vagrant configurations for getting started with
Zookeeper and Go.

Please see the [Getting Started with Zookeeper and Go](https://mmcgrana.github.io/.../getting-started-with-zookeeper-and-go.html)
blog post for details on this code and using Zookeeper with Go.

The basic flow is to install a recent version of Vagrant with
Virtualbox and then:

```console
$ vagrant up
$ vagrant ssh go
(go) ~ $ go get github.com/samuel/go-zookeeper/zk
(go) ~ $ go get github.com/mmcgrana/zk
(go) ~ $ go run /vagrant/0-ping.go
```

### Todo

failure simulation: no zookeper running
failure simulation: kill minority of servers
failure simulation: kill majority of servers
failure simulation: restore servers from backup
zookeeper acls?
jespen?
zk library documentation pull requests
blog post draft
blog post peer review
blog post publication
blog post marketing
