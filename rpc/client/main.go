package main

import (
	"fmt"
	"github.com/fisher310/goder/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{A: 10, B: 3}, &result)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	err = client.Call("DemoService.Div", rpcdemo.Args{A: 10, B: 0}, &result)
	if err != nil {
		// print the error
		fmt.Println(err)
	} else {
		// add in hotfix1 again
		// add some in distribute branch
		// add some in hotfix 2
		fmt.Println(result)
	}
}
