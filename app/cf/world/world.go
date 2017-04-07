package world

import (
	"fmt"
	"net/http"

	// https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/examples/storage/appengine/app.go
	"cloud.google.com/go/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/_ah/warmup", warmupHandler)
	http.HandleFunc("/", handler)
	http.HandleFunc("/secret/", secret)
}

func secret(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		code := http.StatusUnauthorized
		http.Error(w, http.StatusText(code), code)
		return
	}

	log.Infof(ctx, "User: %#v", u)

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>World %s!</h1> <p>this is a secret XX page</p>`, u)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	u := user.Current(ctx)
	log.Infof(ctx, "Unders world service %s", u)
	fmt.Fprint(w, "World service v1! (Standard Environment)")

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
