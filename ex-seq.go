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

	flags := int32(zk.FlagSequence | zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)

	for i:=0; i<3; i++ {
		path, err := conn.Create("/seq-", []byte("data"), flags, acl)
		must(err)
		fmt.Printf("%d %s\n", i, path)
	}
}
