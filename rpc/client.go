package main

import (
	"log"
	"net/rpc/jsonrpc"
)

// Args 外部から呼ばれるときの引数
type Args struct {
	A, B int
}

func main() {
	client, _ := jsonrpc.Dial("tcp", "localhost:18888")

	var result int
	args := &Args{4, 5}
	err := client.Call("Calculator.Multiply", args, &result)
	if err != nil {
		panic(err)
	}
	log.Printf("4 x 5 = %d\n", result)
}
