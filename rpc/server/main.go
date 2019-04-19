package main

import (
	"github.com/fisher310/goder/rpc"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {

	err := rpc.Register(rpcdemo.DemoService{})
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}

	log.Println("log started at port 8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

}
