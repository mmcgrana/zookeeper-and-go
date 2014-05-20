package main

import (
	"fmt"
	"time"
	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	conn1, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	acl := zk.WorldACL(zk.PermAll)
	_, err = conn1.Create("/ephemeral", []byte("here"), zk.FlagEphemeral, acl)
	if err != nil {
		panic(err)
	}

	conn2, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}

	exists, _, err := conn2.Exists("/ephemeral")
	if err != nil {
		panic(err)
	}
	fmt.Printf("before disconnect: %+v\n", exists)

	conn1.Close()
	time.Sleep(time.Second * 2)

	exists, _, err = conn2.Exists("/ephemeral")
	if err != nil {
		panic(err)
	}
	fmt.Printf("after disconnect: %+v\n", exists)
}
