package main

import (
	"log"
	"net/http"
	"remy/internal/database"
	"remy/internal/server"
)

func main() {
	log.Println("Starting....")
	go server.RunGRPC()
	if e := database.Init(); e != nil {
		panic(e)
	}
	http.ListenAndServe("0.0.0.0:8080", nil)
}
