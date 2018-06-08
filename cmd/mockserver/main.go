package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yedamao/go_smgp/smgp/smgptest"
)

var (
	addr = flag.String("addr", ":8890", "addr(本地监听地址)")
)

func init() {
	flag.Parse()
}

func main() {
	server, err := smgptest.NewServer(*addr)
	if err != nil {
		flag.Usage()
		os.Exit(-1)
	}

	HandleSignals(server.Stop)

	server.Run()

	fmt.Println("Done")
}
