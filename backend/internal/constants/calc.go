package constants

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func CalcTime(distance float64) time.Duration {
	var d time.Duration = time.Duration(int(distance / float64(AvgVelo)))
	return d
}

func GetTime(p1, p2 *Point) time.Duration {
	resp, err := http.Get(fmt.Sprintf("http://router.project-osrm.org/route/v1/driving/%f,%f;%f,%f?overview=false", p1.X, p1.Y, p2.X, p2.Y))
	if err != nil {
		log.Printf("Get http: %s", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read: %s", err)
	}
	var t RouteResp
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Printf("Unmarshal: %s", err)
	}
	if t.Code == "Ok" {
		log.Println(t.Routes[0].Distance)
		return CalcTime(t.Routes[0].Distance)
	} else {
		return time.Duration(-1)
	}
}
