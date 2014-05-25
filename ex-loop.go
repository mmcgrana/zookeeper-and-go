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

	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	fmt.Println("create")
	_, err := conn.Create("/loop", []byte("here"), flags, acl)
	must(err)

	i := 0
	for {
		fmt.Println("get")
		_, stat, err := conn.Get("/loop")
		must(err)
		fmt.Println("set")
		_, err = conn.Set("/loop", []byte("here"), stat.Version)
		must(err)
		time.Sleep(time.Second)
		i++
	}
}
