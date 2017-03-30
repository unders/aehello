package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/unders/aehello/app/pkg/health"
	"github.com/unders/aehello/pkg/signal"
)

var (
	// Version the version of the service.
	Version = "No Version Provided"

	// Buildstamp the time the service was built.
	Buildstamp = "No Buildstamp provided"

	// Githash the git commit hash
	Githash = "No Githash provided"
)

// addr is the default address for appengine services
var addr = "0.0.0.0:8080"

// log messages
const (
	started          = "info=started app=helloworld\n"
	release          = "info=version:%s buildstamp:%s githash:%s app=helloworld\n"
	stopped          = "info=stopped app=helloworld\n"
	stoppedWithError = "err=%s app=helloworld\n"
	running          = "info=listens on address %s app=helloworld\n"
	gotStopSignal    = "\ninfo=got signal %s app=helloworld\n"
	request          = "info=%s %s  [ 200 ok ] app=helloworld\n"
	notFound         = "info=%s %s  [ 404 not found ] app=helloworld\n"
)

func main() {
	if ok := run(os.Args, os.Stdout); !ok {
		os.Exit(1)
	}
}

// ENV
// GCLOUD_PROJECT
// GAE_INSTANCE
// GAE_SERVICE
// GAE_VERSION

// The following HTTP headers are now included with all requests:
// X-FORWARDED-FOR
// X-CLOUD-TRACE-CONTEXT
// X-FORWARDED-PROTO

// some requests
// Your application should handle the special country code ZZ (unknown country).
// X-AppEngine-Country        # ISO 3166-1 alpha-2 country code

// X-AppEngine-Region

// Headers to set: Strict-Transport-Security: max-age=31536000; includeSubDomains

func run(args []string, log io.Writer) bool {
	fmt.Fprint(log, started)
	fmt.Fprintf(log, release, Version, Buildstamp, Githash)
	httpAddr := parseAddr(args)

	http.HandleFunc(health.Handler())
	http.HandleFunc("/", landing(log))

	ch := make(chan error, 1)
	go func() {
		ch <- http.ListenAndServe(httpAddr, nil)
	}()

	fmt.Fprintf(log, running, httpAddr)

	select {
	case err := <-ch:
		fmt.Fprintf(log, stoppedWithError, err)
		fmt.Fprint(log, stopped)
		return false
	case sig := <-signal.Wait():
		fmt.Fprintf(log, gotStopSignal, sig)
		fmt.Fprint(log, stopped)
		return true
	}
}

func landing(log io.Writer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			fmt.Fprintf(log, notFound, r.Method, r.RequestURI)
			return
		}
		fmt.Fprint(w, landingPage)
		fmt.Fprintf(log, request, r.Method, r.RequestURI)
	}
}

const landingPage = "Hello World\n"

func parseAddr(args []string) string {
	flag.StringVar(&addr, "http", addr, "HTTP service address.")
	flag.Parse()
	return addr
}
