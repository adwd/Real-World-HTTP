package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Calculator メソッドが属す構造体
type Calculator int

// Multiply RPCで外部から呼ばれるメソッド
func (c *Calculator) Multiply(args Args, result *int) error {
	log.Printf("Multiply called: %d, %d\n", args.A, args.B)
	*result = args.A * args.B
	return nil
}

// Args 外部から呼ばれるときの引数
type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	server := rpc.NewServer()
	server.Register(calculator)
	http.Handle(rpc.DefaultRPCPath, server)
	log.Println("start http listening :18888")
	listener, _ := net.Listen("tcp", ":18888")

	for {
		conn, _ := listener.Accept()
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
