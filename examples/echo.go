package main

import (
	"flag"
	"github.com/kentaro/gyo"
	"log"
)

var apiToken = flag.String("token", "", "API Token")
var port = flag.Int("port", 8080, "Port of the server")
var path = flag.String("path", "/callback", "Callback URL of the server")

func init() {
	flag.Parse()
}

func main() {
	if *apiToken == "" {
		log.Fatalln("API token is required")
	}

	gyo := gyo.NewGyo(*apiToken)
	gyo.Server(*path, *port, func(username string) {
		gyo.Yo(username)
		log.Printf("Sent Yo to %s\n", username)
	})

	return
}
