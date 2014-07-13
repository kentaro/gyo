package main

import (
	"flag"
	"github.com/kentaro/gyo"
	"log"
)

var apiToken = flag.String("token", "", "API Token")
var port = flag.Int("port", 8080, "Port of the server")

func init() {
	flag.Parse()
}

func main() {
	gyo := gyo.NewGyo(*apiToken)
	gyo.Server("/callback", *port, func(username string) {
		gyo.Yo(username)
		log.Printf("Sent Yo to %s\n", username)
	})

	return
}
