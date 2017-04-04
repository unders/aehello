package mux

import (
	"net/http"

	"github.com/unders/aehello/app/helloworld/log"
	"github.com/unders/aehello/app/pkg/health"
	"github.com/unders/aehello/app/pkg/secure"
)

// New returns http.Handler
//
// Usage :
//         m := mux.New("example.com")
//
func New(host string) http.Handler {
	m := http.NewServeMux()
	m.HandleFunc(health.Handler())
	m.HandleFunc("/", landing)

	return &secure.Handler{
		Mux:        m,
		RedirectTo: "https://" + host + ":443",
	}
}

func landing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if _, err := w.Write(landingPage); err != nil {
		log.Error(err)
	}
}

var landingPage = []byte("Welcome to Hellow world\n")
