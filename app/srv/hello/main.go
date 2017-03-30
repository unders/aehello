package main

import (
	"log"
	"net/http"
	// "github.com/pkg/errors"
)

// Sample helloworld is a basic App Engine flexible app.

import (
	"fmt"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/_ah/health", health)
	log.Print("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello world!")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}
