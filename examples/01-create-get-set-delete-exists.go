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
	servers := strings.Split(os.Getenv("ZOOKEEPER_SERVERS"), ",")
	conn, _, err := zk.Connect(servers, time.Second)
	must(err)
	return conn
}

func main() {
	conn := connect()
	flags := int32(0)
	acl := zk.WorldACL(zk.PermAll)

	path, err := conn.Create("/testkey", []byte("testdata"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n", path)

	data, stat, err := conn.Get("/testkey")
	must(err)
	fmt.Printf("get:    %+v %+v\n", string(data), stat)

	stat, err = conn.Set("/testkey", []byte("newtestdata"), stat.Version)
	must(err)
	fmt.Printf("set:    %+v\n", stat)

	err = conn.Delete("/testkey", stat.Version)
	must(err)
	fmt.Printf("delete: ok\n")

	exists, stat, err := conn.Exists("/testkey")
	must(err)
	fmt.Printf("exists: %+v %+v\n", exists, stat)
}
