package main

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn1, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	found, stat, ech, err := conn1.ExistsW("/check")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n%+v\n%+v\n", found, stat, ech)

	conn2, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(time.Second * 3)
		flags := int32(0)
		acl := zk.WorldACL(zk.PermAll)
		_, err = conn2.Create("/check", []byte("here"), flags, acl)
		if err != nil {
			panic(err)
		}
	}()

	evt := <- ech
	fmt.Printf("%+v\n", evt)

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	_, err = conn1.Create("/check/childa", []byte("here"), flags, acl)
	if err != nil {
		panic(err)
	}

	children, stat, ech, err := conn1.ChildrenW("/check")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n%+v\n%+v\n", children, stat, ech)

	go func() {
		time.Sleep(time.Second * 3)
		flags := int32(0)
		acl := zk.WorldACL(zk.PermAll)
		_, err = conn2.Create("/check/childb", []byte("here"), flags, acl)
		if err != nil {
			panic(err)
		}
	}()

	evt = <- ech
	fmt.Printf("%+v\n", evt)
}
