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

	fmt.Println("presence.create")
	acl := zk.WorldACL(zk.PermAll)
	_, err := conn.Create("/presence", []byte("here"), zk.FlagEphemeral, acl)
	must(err)

	<-make(chan bool)
}
