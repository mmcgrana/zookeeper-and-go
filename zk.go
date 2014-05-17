package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: zk <create|set|get|stat|delete|children> [args]\n")
	os.Exit(1)
}

func connect() *zk.Conn {
	conn, _, err := zk.Connect([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	return conn
}

func input() []byte {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	return data
}

func formatTime(millis int64) string {
	t := time.Unix(0, millis * 1000000)
	return t.Format(time.RFC3339)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		usage()
	}
	subcommand := args[0]
	subcommandArgs := args[1:]
	switch subcommand {
	case "create":
		if len(subcommandArgs) != 1 {
			fmt.Fprintf(os.Stderr, "Usage: zk create <path>\n")
		}
		path := subcommandArgs[0]
		data := input() 
		conn := connect()
		flags := int32(0)
		acl := zk.WorldACL(zk.PermAll)
		_, err := conn.Create(path, data, flags, acl)
		if err != nil {
			panic(err)
		}
	case "set":
		if !(len(subcommandArgs) == 1 || len(subcommandArgs) == 2) {
			fmt.Fprintf(os.Stderr, "Usage: zk set <path> [version]")
		}
		path := subcommandArgs[0]
		readVersion := len(subcommandArgs) == 1
		data := input()
		conn := connect()
		var version int32
		if readVersion {
			_, stat, err := conn.Get(path)
			if err != nil {
				panic(err)
			}
			version = stat.Version
		} else {
			versionParsed, err := strconv.Atoi(subcommandArgs[1])
			if err != nil {
				panic(err)
			}
			version = int32(versionParsed)
		}
		_, err := conn.Set(path, data, version)
		if err != nil {
			panic(err)
		}
	case "get":
		if !(len(subcommandArgs) == 1 || len(subcommandArgs) == 2) {
			fmt.Fprintf(os.Stderr, "Usage: zk get [--watch] <path>")
		}
		conn := connect()
		if len(subcommandArgs) == 1 {
			path := subcommandArgs[0]
			data, _, err := conn.Get(path)
			_, err = os.Stdout.Write(data)
			if err != nil {
				panic(err)
			}
		} else {
			path := subcommandArgs[1]
			data, _, events, err := conn.GetW(path)
			_, err = os.Stdout.Write(data)
			if err != nil {
				panic(err)
			}
			evt := <- events
			if evt.Err != nil {
				panic(evt)
			}
		}
	case "stat":
		if !(len(subcommandArgs) == 1) {
			fmt.Fprintf(os.Stderr, "Usage: zk stat <path>")
		}
		path := subcommandArgs[0]
		conn := connect()
		_, stat, err := conn.Get(path)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(os.Stdout, "Czxid:          %d\n", stat.Czxid)
		fmt.Fprintf(os.Stdout, "Mzxid:          %d\n", stat.Mzxid)
		fmt.Fprintf(os.Stdout, "Ctime:          %s\n", formatTime(stat.Ctime))
		fmt.Fprintf(os.Stdout, "Mtime:          %s\n", formatTime(stat.Mtime))
		fmt.Fprintf(os.Stdout, "Version:        %d\n", stat.Version)
		fmt.Fprintf(os.Stdout, "Cversion:       %d\n", stat.Cversion)
		fmt.Fprintf(os.Stdout, "Aversion:       %d\n", stat.Aversion)
		fmt.Fprintf(os.Stdout, "EphemeralOwner: %d\n", stat.EphemeralOwner)
		fmt.Fprintf(os.Stdout, "DataLength:     %d\n", stat.DataLength)
		fmt.Fprintf(os.Stdout, "Pzxid:          %d\n", stat.Pzxid)
	case "delete":
		if !(len(subcommandArgs) == 1 || len(subcommandArgs) == 2)  {
			fmt.Fprintf(os.Stderr, "Usage: zk delete <path> [version]\n")
		}
		path := subcommandArgs[0]
		readVersion := len(subcommandArgs) == 1
		conn := connect()
		var version int32
		if readVersion {
			_, stat, err := conn.Get(path)
			if err != nil {
				panic(err)
			}
			version = stat.Version
		} else {
			versionParsed, err := strconv.Atoi(subcommandArgs[1])
			if err != nil {
				panic(err)
			}
			version = int32(versionParsed)
		}
		err := conn.Delete(path, version)
		if err != nil {
			panic(err)
		}
	case "children":
		if !(len(subcommandArgs) == 1 || len(subcommandArgs) == 2) {
			fmt.Fprintf(os.Stderr, "Usage: zk children [--watch] <path>")
		}
		conn := connect()
		if len(subcommandArgs) == 1 {
			path := subcommandArgs[0]
			children, _, err := conn.Children(path)
			if err != nil {
				panic(err)
			}
			sort.Strings(children)
			for _, child := range children {
				fmt.Fprintln(os.Stdout, child)
			}
		} else {
			path := subcommandArgs[1]
			children, _, events, err := conn.ChildrenW(path)
			if err != nil {
				panic(err)
			}
			sort.Strings(children)
			for _, child := range children {
				fmt.Fprintln(os.Stdout, child)
			}
			evt := <- events
			if evt.Err != nil {
				panic(evt.Err)
			}
		}
	default:
		fmt.Fprintf(os.Stderr, "Unrecognized subcommand '%s'\n", subcommand)
	}
}
