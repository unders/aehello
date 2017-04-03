package mux

import (
	"fmt"
	"net/http"

	"github.com/unders/aehello/app/helloworld/log"
	"github.com/unders/aehello/app/pkg/health"
)

func New() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc(health.Handler())
	m.HandleFunc("/", landing)
	return m
}

func landing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// FIXME: this will be removed
	xproto := r.Header.Get("X-FORWARDED-PROTO")
	msg := fmt.Sprintf("X-FORWARDED-PROTO:%s", xproto)
	log.Info(msg)

	if _, err := fmt.Fprint(w, landingPage); err != nil {
		log.Error(err)
	}
}

const landingPage = "Welcome to Hellow world\n"
