package main

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	path, err := conn.Create("/testkey", []byte("testdata"), flags, acl)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", path)

	data, stat, err := conn.Get("/testkey")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n%+v\n", string(data), stat)

	stat, err = conn.Set("/testkey", []byte("newtestdata"), stat.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", stat)

	err = conn.Delete("/testkey", stat.Version)
	if err != nil {
		panic(err)
	}

	exists, stat, err := conn.Exists("/testkey")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n%+v\n", exists, stat)
}
