package main

import (
	"fmt"
	"net/http"

	// https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/examples/storage/appengine/app.go
	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

// var out *logging.Logger

func init() {
	http.HandleFunc("/_ah/warmup", warmupHandler)
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", welcome)
	http.HandleFunc("/secret/", secret)
}

func secret(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	log.Infof(ctx, "User: %#v", u)

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Hello %s!</h1> <p>this is a secret page</p>`, u)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}

	log.Infof(ctx, "User: %#v", u)
	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	log.Infof(ctx, "Unders Request: %s", r.RequestURI)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, world 11! (Standard Environment)")

	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
		return
	}

	if err := client.Close(); err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
		return
	}

}

func warmupHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	// Perform warmup tasks, including ones that require a context,
	// such as retrieving data from Datastore.

	log.Infof(ctx, "Warmup: %s", r.RequestURI)
}
