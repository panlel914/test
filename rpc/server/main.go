package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"test/rpc"
)

func main() {
	rpc.Register(rpcdemo.RpcDemoServer{})
	listener, err := net.Listen("tcp",":4444")
	if err != nil{
		panic(err)
	}

	for{
		coon, err := listener.Accept()
		if err != nil{
			log.Printf("err: %v", err)
		}

		go jsonrpc.ServeConn(coon)
	}
}
