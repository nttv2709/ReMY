package constants

import "flag"

var (
	Port    = flag.Int("port", 50051, "The server port")
	KmH2MM  = 16.67
	AvgVelo = 35 * KmH2MM // km/h
)
