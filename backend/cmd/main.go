package main

import (
	"sample_shecodes2022/internal/database"
	"sample_shecodes2022/internal/server"
)

func main() {
	go server.RunGRPC()
	database.Init()
}
