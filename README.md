## Zookeeper and Go

Example code and Vagrant configurations for getting started with
Zookeeper and Go.

Please see the [Getting Started with Zookeeper and Go](https://mmcgrana.github.io/2014/05/getting-started-with-zookeeper-and-go.html)
blog post for details on this code and using Zookeeper with Go.

The basic flow is to install a recent version of Vagrant with
VirtualBox and then:

```console
$ vagrant up
$ vagrant ssh go
(go) ~ $ go get github.com/samuel/go-zookeeper/zk
(go) ~ $ go get github.com/mmcgrana/zk
(go) ~ $ go run /vagrant/ex-ping.go
```

`ex-*.go` files are example programs running with `go run` as above.

`vm-*.sh` are VM configuration files, used in `vagrant up` above.

`sim-*.txt` are failure simulation notes, showing how to run basic
failure simulations in the environment described above and the
results they should produce.
