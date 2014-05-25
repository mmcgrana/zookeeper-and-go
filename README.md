## Zookeeper and Go

Source code and Vagrant configurations for getting started with
Zookeeper and Go.

Please see the [Getting Started with Zookeeper and Go](https://mmcgrana.github.io/.../getting-started-with-zookeeper-and-go.html)
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

In addition to the content covered in the above blog post, this
project also contains notes on basic failure simulations conducted
using the Zookeeper/Go environment described here; see files
`sim-*.sh`.

### Todo

failure simulation: kill minority of servers
  -> everything works, some weird logging?
failure simulation: kill majority of servers
  -> can read, can't write?
failure simulation: restart cluster
  -> how does it even work
failure simulation: network hang
  -> ???
zk library documentation pull requests
review zk book
blog peer review
blog publication
blog marketing
