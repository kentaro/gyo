package gyo

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const APIBASE = "http://api.justyo.co"

type Gyo struct {
	APIToken string
}

func NewGyo(token string) (self *Gyo) {
	self = &Gyo{APIToken: token}
	return
}

func (self *Gyo) YoAll() (res *http.Response, err error) {
	res, err = self.dispatchRequest("/yoall/", map[string]string{})
	return
}

func (self *Gyo) Yo(username string) (res *http.Response, err error) {
	res, err = self.dispatchRequest("/yo/", map[string]string{"username": username})
	return
}

func (self *Gyo) Server(path string, port int, callback func(username string)) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		callback(r.URL.Query().Get("username"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func (self *Gyo) dispatchRequest(path string, params map[string]string) (res *http.Response, err error) {
	values := url.Values{}
	values.Set("api_token", self.APIToken)

	for key, value := range params {
		values.Set(key, value)
	}

	res, err = http.PostForm(APIBASE+path, values)
	return
}
