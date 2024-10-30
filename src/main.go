package main 

import (
	"fmt"
	"net/http"
	// "time"
)

type URLMap struct {
	urls map[string]string
}

func (um *URLMap) Shorten (writer http.ResponseWriter, req *http.Request) {

}

func (um *URLMap) Redirected (writer http.ResponseWriter, req *http.Request) {
	
	shortKey := req.URL.Path[len("/shortened/"):]
	if shortKey == "" {
		http.Error(writer, "The shortened key is missing", http.StatusBadRequest)
		return
	}	

	longURL, found := um.urls[shortKey]
	if !found {
		http.Error(writer, "The shortened key wasn't fount", http.StatusNotFound)
		return
	}

	http.Redirect(writer, req, longURL, http.StatusMovedPermanently)
}

func main () {

	shortener := &URLMap {
		urls: make(map[string]string),
	}
	
	http.HandleFunc("/link", shortener.Shorten)
	http.HandleFunc("/shortened/", shortener.Redirected)	

	fmt.Println("Started server on port 8080.")
	http.ListenAndServe(":8080", nil)

}
