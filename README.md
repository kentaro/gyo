# Gyo - Yo for Go

Gyo is a Golang interface to [Yo API](http://developer.justyo.co/yo/docs.html).

## Usage

Here's an echo server example using Yo API.

```go
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
```

## Author

* [Kentaro Kuribayashi](http://kentarok.org)

## License

* [MIT](http://kentaro.mit-license.org/)
