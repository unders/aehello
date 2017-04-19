package main

import (
	"fmt"
	"net/http"

	// https://github.com/GoogleCloudPlatform/google-cloud-go/blob/master/examples/storage/appengine/app.go
	"cloud.google.com/go/storage"
	"github.com/unders/aehello/app/pkg/jwt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
)

func init() {
	http.HandleFunc("/_ah/warmup", warmupHandler)
	http.HandleFunc("/", handler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/secret/", secret)
	http.HandleFunc("/token/", token)
	http.HandleFunc("/read-token/", getToken)
}

// https://github.com/drichardson/appengine/blob/master/signedrequest/request.go
// https://github.com/drichardson/appengine/tree/master/signature
// https://github.com/someone1/gcp-jwt-go
// https://github.com/someone1/gcp-jwt-go/blob/master/appengine.go
func token(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	err := jwt.SetShortToken(ctx, w, jwt.User{Name: "Unders", Email: "unders@example.com"})
	if err != nil {
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<h1>Generating token failed</h1> <p>Error: %s</p>`, err)
		return
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	t := w.Header().Get(jwt.HeaderKey)
	fmt.Fprintf(w, `<h1>Token is set</h1> <p> %s</p>`, t)
}

func getToken(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	user, err := jwt.GetUser(ctx, r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Header().Set("Content-type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `<h1>Read token failed</h1> <p>Error: %s</p>`, err)
		return
	}

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Read token. User:</h1> <p> %s</p>`, user)
}

func secret(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)

	log.Infof(ctx, "User: %#v", u)

	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>Hello %s!</h1> <p>this is a secret page</p>`, u)
}

func login(w http.ResponseWriter, r *http.Request) {
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
