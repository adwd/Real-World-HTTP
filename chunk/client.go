package main

import "net/http"
import "bufio"
import "io"
import "log"
import "bytes"

func main() {
	resp, _ := http.Get("http://localhost:18888/chunk")
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		log.Println(string(bytes.TrimSpace(line)))
	}
}
