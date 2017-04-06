package main

import (
	"fmt"
	"net/http"

	// https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/examples/storage/appengine/app.go
	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

// var out *logging.Logger

func init() {
	http.HandleFunc("/_ah/warmup", warmupHandler)
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	log.Infof(ctx, "Unders Request: %s", r.RequestURI)
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
