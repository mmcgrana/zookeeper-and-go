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

	path, err := conn.Create("/01", []byte("data"), flags, acl)
	must(err)
	fmt.Printf("create: %+v\n", path)

	data, stat, err := conn.Get("/01")
	must(err)
	fmt.Printf("get:    %+v %+v\n", string(data), stat)

	stat, err = conn.Set("/01", []byte("newdata"), stat.Version)
	must(err)
	fmt.Printf("set:    %+v\n", stat)

	err = conn.Delete("/01", stat.Version)
	must(err)
	fmt.Printf("delete: ok\n")

	exists, stat, err := conn.Exists("/01")
	must(err)
	fmt.Printf("exists: %+v %+v\n", exists, stat)
}
