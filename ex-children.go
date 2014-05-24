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
	conn := connect()
	defer conn.Close()

	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)
	_, err = conn.Create("/namespace", []byte(""), flags, acl)
	must(err)

	_, err = conn.Create("/namespace/nested", []byte(""), flags, acl)
	must(err)
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("/namespace/key%d", i)
		data := []byte(fmt.Sprintf("data%d", i))
		path, err := conn.Create(key, data, flags, acl)
		must(err)
		fmt.Printf("%+v\n", path)
	}
	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("/namespace/nested/key%d", i)
		data := []byte(fmt.Sprintf("nesteddata%d", i))
		path, err := conn.Create(key, data, flags, acl)
		must(err)
		fmt.Printf("%+v\n", path)
	}
	fmt.Println()

	fmt.Println("zk.children")
	children, _, err := conn.Children("/namespace")
	must(err)
	sort.Strings(children)
	for _, path := range children {
		_, stat, err := conn.Get("/namespace/" + path)
		must(err)
		fmt.Printf("%+v %d\n", path, stat.NumChildren)
	}
	children, _, err = conn.Children("/namespace/nested")
	must(err)
	sort.Strings(children)
	for _, path := range children {
		fmt.Printf("%+v\n", path)
	}
}
