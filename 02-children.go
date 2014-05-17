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
	_, err = conn.Create("/namespace", []byte(""), flags, acl)
	if err != nil {
		panic(err)
	}
	_, err = conn.Create("/namespace/nested", []byte(""), flags, acl)
	if err != nil {
		panic(err)
	}
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("/namespace/key%d", i)
		data := []byte(fmt.Sprintf("data%d", i))
		path, err := conn.Create(key, data, flags, acl)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", path)
	}
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("/namespace/nested/key%d", i)
		data := []byte(fmt.Sprintf("nesteddata%d", i))
		path, err := conn.Create(key, data, flags, acl)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", path)
	}
	fmt.Println()

	fmt.Println("zk.children")
	children, _, err := conn.Children("/namespace")
	if err != nil {
		panic(err)
	}
	sort.Strings(children)
	for _, path := range children {
		_, stat, err := conn.Get("/namespace/" + path)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v %d\n", path, stat.NumChildren)
	}
	children, _, err = conn.Children("/namespace/nested")
	if err != nil {
		panic(err)
	}
	sort.Strings(children)
	for _, path := range children {
		fmt.Printf("%+v\n", path)
	}
}
