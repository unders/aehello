package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/unders/aehello/app/helloworld/config"
	"github.com/unders/aehello/app/helloworld/log"
	"github.com/unders/aehello/app/pkg/health"
	"github.com/unders/aehello/pkg/signal"
)

// These are set from the build script
var (
	Version    = "The version of the service"
	Buildstamp = "the time it was built"
	Githash    = "The git commit hash"
)

func main() {
	log.Started()
	log.Release(Version, Buildstamp, Githash)

	o, err := config.New(os.Args)
	log.Config(fmt.Sprintf("%#v", o))
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	if ok := run(o); !ok {
		os.Exit(1)
	}
}

// Headers to set: Strict-Transport-Security: max-age=31536000; includeSubDomains

func run(o config.Options) bool {
	l := log.Init(o)
	defer l.Close()

	http.HandleFunc(health.Handler())
	http.HandleFunc("/", landing)

	ch := make(chan error, 1)
	go func() {
		ch <- http.ListenAndServe(o.Server.Addr, nil)
	}()

	log.Running(o.Server.Addr)

	select {
	case err := <-ch:
		log.Error(err)
		log.Stopped()
		return false
	case sig := <-signal.Wait():
		log.GotStopSignal(sig)
		log.Stopped()
		return true
	}
}

func landing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/error" {
		log.Error(errors.New("This is an test error"))
	}

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
