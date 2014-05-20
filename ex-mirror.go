package main

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func mirror(conn *zk.Conn, path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)
	go func() {
		for {
			snapshot, _, events, err := conn.ChildrenW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- snapshot
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
		}
	}()
	return snapshots, errors
}

func main() {
	conn1, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	acl := zk.WorldACL(zk.PermAll)
	_, err = conn1.Create("/mirror", []byte("here"), 0, acl)
	if err != nil {
		panic(err)
	}

	snapshots, errors := mirror(conn1, "/mirror")
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				fmt.Printf("%+v\n", snapshot)
			case err := <-errors:
				panic(err)
			}
		}
	}()

	conn2, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	_, err = conn2.Create("/mirror/one", []byte("one"), zk.FlagEphemeral, acl)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	_, err = conn2.Create("/mirror/two", []byte("two"), zk.FlagEphemeral, acl)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	_, err = conn2.Set("/mirror/one", []byte("one new"), 0)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	err = conn2.Delete("/mirror/two", 0)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
	_, err = conn2.Create("/mirror/three", []byte("three"), zk.FlagEphemeral, acl)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second)
}
