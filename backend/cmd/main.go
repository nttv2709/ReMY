package main

import (
	"log"
	"sample_shecodes2022/internal/database"
	"sample_shecodes2022/internal/server"
)

func main() {
	log.Println("Starting....")
	go server.RunGRPC()
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
}
