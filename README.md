## Zookeeper and Go

Example Go code, Vagrant configurations, and Dockerfiles for getting
started with Zookeeper and Go.

Based on the post [Getting Started with Zookeeper and Go](https://mmcgrana.github.io/2014/05/getting-started-with-zookeeper-and-go.html),
but modified to use Docker in addition to Vagrant.

### Usage

First, ensure you have a recent version of Vagrant with Virtualbox
installed. Then run:

```console
$ vagrant up
$ vagrant ssh
$ sudo docker run --name go.1 -i -t -v /vagrant:/vagrant --env-file /etc/go.env go
```

In a separate terminal:

```console
$ vagrant ssh
$ sudo pipework br1 go.1 "$GO_IP/24"
```

Then back in the first:

```console
$ go run /vagrant/ex-ping.go
```

This should write "ok" if everything is working.

Run other `ex-*.go` example programs with `go run` as above.

`sim-*.txt` are failure simulation notes, showing how to run basic
failure simulations in this environment and the results they should
produce.

### Vagrant and Docker Configuration

This system uses Vagrant and Docker to build a complete local
development environment. See the `Vagrantfile` for details.

Vagrant provides a single Ubuntu VM in which we run 3 containers
with Zookeepers 1 container with a Go environment. All containers
are connected by a shared bridge with networking separate from the
host's.
