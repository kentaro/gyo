package main

import (
	"flag"
	"github.com/kentaro/gyo"
	"log"
)

var apiToken = flag.String("token", "", "API Token")
var userName = flag.String("username", "", "The user to whom you want to Yo")

func init() {
	flag.Parse()
}

func main() {
	gyo := gyo.NewGyo(*apiToken)
	var err error

	switch *userName {
	case ":all":
		_, err = gyo.YoAll()
	default:
		_, err = gyo.Yo(*userName)
	}

	if err != nil {
		log.Fatal(err)
	}

	return
}
