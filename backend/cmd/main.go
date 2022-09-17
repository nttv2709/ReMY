package main

import (
	"log"
	"net/http"
	"remy/internal/server"
)

func main() {
	log.Println("Starting....")
	go server.RunGRPC()
	http.ListenAndServe("0.0.0.0:8080", nil)
}
