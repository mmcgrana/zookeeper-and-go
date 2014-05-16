package main

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	fmt.Println("zk.connect")
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Println("zk.children")
	children, stat, err := conn.Children("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)

	// fmt.Println("zk.create")
	// acl := zk.WorldACL(zk.PermAll)
	// info, err := conn.Create("/test", []byte("testdata"), 0, acl)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%+v\n", info)
	
	fmt.Println("zk.get")
	data, stat, err := conn.Get("/test")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", string(data), stat)

	fmt.Println("zk.set")
	stat, err = conn.Set("/test", []byte("newtestdata"), stat.Version)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", stat)

	fmt.Println("zk.children")
	children, stat, err = conn.Children("/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)
}
