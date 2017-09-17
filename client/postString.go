package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("テキスト")
	resp, _ := http.Post("http://localhost:18888", "text/plain", reader)
	log.Println("status", resp.Status)
}
