package main

import "net/http"
import "io/ioutil"
import "log"

func main() {
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
	log.Println("status code:", resp.StatusCode)
	// log.Println("headers", resp.Header)
	log.Println("headers")
	for key, value := range resp.Header {
		log.Println(key, value)
	}
	log.Println(string(body))
}
