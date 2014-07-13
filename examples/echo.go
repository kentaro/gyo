package main

import (
	"flag"
	"github.com/kentaro/gyo"
	"log"
)

var apiToken = flag.String("token", "", "API Token")

func init() {
	flag.Parse()
}

func main() {
	gyo := gyo.NewGyo(*apiToken)
	gyo.Server("/callback", 8080, func(username string) {
		gyo.Yo(username)
		log.Printf("Sent Yo to %s\n", username)
	})

	return
}
