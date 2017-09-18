package main

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	// open TCP socket
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, _ := dialer.Dial("tcp", "localhost:18888")
	defer conn.Close()

	// send request
	request, _ := http.NewRequest("GET", "http://localhost:18888/chunk", nil)
	request.Write(conn)

	// read
	reader := bufio.NewReader(conn)
	// read header
	resp, _ := http.ReadResponse(reader, request)
	if resp.TransferEncoding[0] != "chunked" {
		panic("wrong transfer encoding")
	}

	for {
		// fetch size
		sizeStr, _ := reader.ReadBytes('\n')
		log.Println("size:", string(sizeStr))

		// parse 16 basis size, close if size is 0
		size, _ := strconv.ParseInt(string(sizeStr[:len(sizeStr)-2]), 16, 64)
		if size == 0 {
			break
		}
		// allocate buffer and read
		line := make([]byte, int(size))
		reader.Read(line)
		reader.Discard(2)
		log.Println("  ", string(line))
	}
}
