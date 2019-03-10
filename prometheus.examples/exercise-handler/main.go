package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"math/rand"
	"net/http"
)

var (
	url = flag.String("server-url", "http://localhost:8085", "The URL of the HTTP server.")
)

func main () {
	flag.Parse()

	for {
		_, err := http.Get(fmt.Sprintf("%s/%s", *url, randomEndpoint()))
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Duration(500+rand.Intn(1000))*time.Millisecond)
	}
}

func randomEndpoint() string {
	if rand.Float32() < 0.48 {
		return "hello"
	}
	return "goodbye"
}
