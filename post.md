# Getting started with Zookeeper and Go

In this post we'll show how to use the [Zookeper](http://zookeeper.apache.org/)
distributed storage system from [Go](http://golang.org/). We'll set
up a Zookeeper and Go environment using Vagrant, perform common
Zookeeper operations from the Go client, and conducting failure
simulations to test the resiliency of our system.

## Zookeeper and Go Environment

For the Go client examples described below, we'll need a multi-node
Zookeeper cluster and a working Go development environment. To set
this up we'll use [Vagrant](http://www.vagrantup.com/), a tool to
easily build reproducible development environments.



$ vagrant version
Vagrant 1.5.4


$ vagrant up
...


$ vagrant ssh 

## Feedback Welcome

We welcome feedback on 