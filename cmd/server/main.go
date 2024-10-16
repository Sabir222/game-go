package main

import (
	"flag"
	"github.com/Sabir222/game-go/server"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", server.Server)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
