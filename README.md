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
$ sudo /bin/go-console
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

### Restoring the VM

Failure simulations will leave the VM in an broken state, by
design. It's best to restore by rebuilding from scratch:

```console
$ docker ps -a | tail -n +2 | cut -d ' ' -f 1 | xargs -n 1 docker rm -f
$ /vagrant/vm.sh
```

This operation should be much faster than the initial boot of the
VM.
