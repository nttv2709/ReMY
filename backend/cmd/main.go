package main

import (
	"log"
	"net/http"
	"remy/internal/constants"
	"remy/internal/server"
)

func main() {
	log.Println("Starting....")
	log.Println(constants.GetTime(constants.NewPoint(13.388860, 52.517037), constants.NewPoint(13.397634, 52.529407)))
	go server.RunGRPC()
	http.ListenAndServe("0.0.0.0:8080", nil)
}
