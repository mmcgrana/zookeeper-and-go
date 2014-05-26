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
	conn1 := connect()
	defer conn1.Close()

	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	found, _, ech, err := conn1.ExistsW("/watch")
	must(err)
	fmt.Printf("found: %t\n", found)

	conn2 := connect()
	must(err)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("creating znode")
		_, err = conn2.Create("/watch", []byte("here"), flags, acl)
		must(err)
	}()

	evt := <-ech
	fmt.Println("watch fired")
	must(evt.Err)

	found, _, err = conn1.Exists("/watch")
	must(err)
	fmt.Printf("found: %t\n", found)
}
