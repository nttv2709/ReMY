package main

import (
	"log"
	"net/http"
	"sample_shecodes2022/internal/database"
	"sample_shecodes2022/internal/server"
)

func main() {
	log.Println("Starting....")
	go server.RunGRPC()
	if e := database.Init(); e != nil {
		panic(e)
	}
	http.ListenAndServe("0.0.0.0:8080", nil)
}
