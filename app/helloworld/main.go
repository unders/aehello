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
	started          = "\n** STARTED **\ninfo=started app=helloworld\n"
	release          = "info=version:%s buildstamp:%s githash:%s app=helloworld\n"
	stopped          = "info=stopped app=helloworld\n** STOPPED **\n"
	stoppedWithError = "err=%s app=helloworld\n"
	running          = "info=listens on address %s app=helloworld\n"
	gotStopSignal    = "\ninfo=got signal %s app=helloworld\n"
)

func main() {
	fmt.Println(started)
	// test to log to stderr
	fmt.Fprintln(os.Stderr, "err=this is an error!")
	logEnv()

	if ok := run(os.Args, os.Stdout); !ok {
		os.Exit(1)
	}
}

// Headers to set: Strict-Transport-Security: max-age=31536000; includeSubDomains

func run(args []string, log io.Writer) bool {
	fmt.Fprintf(log, release, Version, Buildstamp, Githash)
	httpAddr := parseAddr(args)

	http.HandleFunc(health.Handler())
	http.HandleFunc("/", landing)

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

func logEnv() {
	proj := os.Getenv("GCLOUD_PROJECT")
	fmt.Println("GCLOUD_PROJECT", proj)

	instance := os.Getenv("GAE_INSTANCE")
	fmt.Println("GAE_INSTANCE", instance)

	srv := os.Getenv("GAE_SERVICE")
	fmt.Println("GAE_SERVICE", srv)

	version := os.Getenv("GAE_VERSION")
	fmt.Println("GAE_VERSION", version)

	my := os.Getenv("MY_ENV")
	fmt.Println("MY_ENV", my)
}

// some requests
// Your application should handle the special country code ZZ (unknown country).
// X-AppEngine-Country        # ISO 3166-1 alpha-2 country code

// X-AppEngine-Region
func landing(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// The following HTTP headers are now included with all requests:
	xfor := r.Header.Get("X-FORWARDED-FOR")
	fmt.Println("X-FORWARDED-FOR", xfor)
	xctx := r.Header.Get("X-CLOUD-TRACE-CONTEXT")
	fmt.Println("X-CLOUD-TRACE-CONTEXT", xctx)
	xproto := r.Header.Get("X-FORWARDED-PROTO")
	fmt.Println("X-FORWARDED-PROTO", xproto)

	fmt.Fprint(w, landingPage)
}

const landingPage = "Welcome to Hellow world\n"

func parseAddr(args []string) string {
	flag.StringVar(&addr, "http", addr, "HTTP service address.")
	flag.Parse()
	return addr
}
