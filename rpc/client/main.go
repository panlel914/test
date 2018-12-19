package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"test/rpc"
)

func main() {
	coon, err := net.Dial("tcp", ":4444")
	if err != nil{
		fmt.Println(err)
	}
	var(
		result *int
	)
	client := jsonrpc.NewClient(coon)
	err = client.Call("RpcDemoServer.Add", rpcdemo.Args{1,5}, &result)
	if err != nil{
		fmt.Println("err:",err)
	}

	fmt.Println("result:", *result)
}
