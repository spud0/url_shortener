package main 

import (
	"fmt"
	"net/http"
	"time"
)

type URLMap struct {
	urls map{string}string
}

func (us *URLMap) Shorten (writer http.ResponseWriter, req *http.Request) {

}

func (us *URLMap) Redirected (writer http.ResponseWriter, req *http.Request) {

}

func main () {
	shortener := &URLMap {
		urls: make(map[string]string),
	}

	http.HandleFunc("/link", shortener.Shorten)
	http.HandleFunc("/shortened/", shortener.Redirected)	
	http.ListenAndServe(":8080", nil)
}
