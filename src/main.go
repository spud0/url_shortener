package main 

import (
	"fmt"
	"net/http"
	"github.com/spud0/url_shortener/internal/db"
	// "github.com/spud0/url_shortener/internal/cache"
)

// Replace with DB connection
// Also add caching to check for 
type URLMap struct {
	urls map[string]string
}

func (um *URLMap) Shorten (writer http.ResponseWriter, req *http.Request) {
		
	if req.Method != http.MethodPost {
		http.Error(writer, "Invalid request method.", http.StatusMethodNotAllowed)
		return
	}	

	longURL := req.FormValue("url") 
	if longURL == "" {
		http.Error(writer, "The URL is missing.", http.StatusBadRequest)
		return
	}

	shortKey := "1"
	um.urls[shortKey] = longURL
	shortened := fmt.Sprintf("http://localhost:8080/shortened/%s", shortKey)

	writer.Header().Set("Content-Type", "text/html")
	responseHTML := fmt.Sprintf(`
		<h2>URL Shortener</h2>
		<p>Original URL: %s</p>
		<p>Shortened URL: <a href="%s">%s</a></p>
		<form method="post" action="/link">
			<input type="text" name="url" placeholder="Enter a URL">
			<input type="submit" value="Shorten">
		</form>
	`, longURL, shortened, shortened)
	fmt.Fprintf(writer, responseHTML)

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

func (um *URLMap) Index (writer http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		writer.Header().Set("Content-Type", "text/html")
		formHTML := `
			<h2>URL Shortener</h2>
			<form method="post" action="/link">
				<input type="text" name="url" placeholder="Enter a URL">
				<input type="submit" value="Shorten">
			</form>
		`
		fmt.Fprint(writer, formHTML)
		return
	}
}

func main () {

	shortener := &URLMap {
		urls: make(map[string]string),
	}
	
	http.HandleFunc("/shortened/", shortener.Redirected)	
	http.HandleFunc("/link", shortener.Shorten)
	http.HandleFunc("/", shortener.Index)

	fmt.Println("Started server on port 8080.")

	conn, err := db.NewDB() 
	if err != nil {
		fmt.Println("[error]: %s", err)
	}

	fmt.Println("Value from package is:", conn)

	http.ListenAndServe(":8080", nil)

}
