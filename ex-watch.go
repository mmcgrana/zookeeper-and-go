package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() *zk.Conn {
	zksStr := os.Getenv("ZOOKEEPER_SERVERS")
	zks := strings.Split(zksStr, ",")
	conn, _, err := zk.Connect(zks, time.Second)
	must(err)
	return conn
}

func main() {
	conn1 := connect()
	defer conn1.Close()

	found, stat, ech, err := conn1.ExistsW("/check")
	must(err)
	fmt.Printf("%+v\n%+v\n%+v\n", found, stat, ech)

	conn2 := connect()
	must(err)

	go func() {
		time.Sleep(time.Second * 3)
		flags := int32(0)
		acl := zk.WorldACL(zk.PermAll)
		_, err = conn2.Create("/check", []byte("here"), flags, acl)
		must(err)
	}()

	evt := <- ech
	fmt.Printf("%+v\n", evt)

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	_, err = conn1.Create("/check/childa", []byte("here"), flags, acl)
	must(err)

	children, stat, ech, err := conn1.ChildrenW("/check")
	must(err)
	fmt.Printf("%+v\n%+v\n%+v\n", children, stat, ech)

	go func() {
		time.Sleep(time.Second * 3)
		flags := int32(0)
		acl := zk.WorldACL(zk.PermAll)
		_, err = conn2.Create("/check/childb", []byte("here"), flags, acl)
		must(err)
	}()

	evt = <- ech
	fmt.Printf("%+v\n", evt)
}
