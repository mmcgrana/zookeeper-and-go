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
	_, _, err := conn.Exists("/ping")
	must(err)
	fmt.Println("ok")
}
