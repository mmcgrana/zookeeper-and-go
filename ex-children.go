package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strings"
	"time"
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

	_, err := conn.Create("/dir", []byte("data-parent"), flags, acl)
	must(err)

	for i := 1; i <= 3; i++ {
		key := fmt.Sprintf("/dir/key%d", i)
		data := []byte(fmt.Sprintf("data-child-%d", i))
		path, err := conn.Create(key, data, flags, acl)
		must(err)
		fmt.Printf("%+v\n", path)
	}

	data, _, err := conn.Get("/dir")
	fmt.Printf("/dir: %s\n", string(data))

	children, _, err := conn.Children("/dir")
	must(err)
	for _, name := range children {
		data, _, err := conn.Get("/dir/" + name)
		must(err)
		fmt.Printf("/dir/%s: %s\n", name, string(data))
		err = conn.Delete("/dir/"+name, 0)
	}

	err = conn.Delete("/dir", 0)
	must(err)
}
