package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/unders/aehello/app/helloworld/config"
	"github.com/unders/aehello/app/helloworld/log"
	"github.com/unders/aehello/app/helloworld/mux"
	"github.com/unders/aehello/app/pkg/http"
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

	if err := run(o); err != nil {
		os.Exit(1)
	}
}

func run(o config.Options) error {
	var err error

	l := log.Init(o)
	defer l.Close()

	s := http.Server{
		Addr:         o.HTTP.Addr,
		ReadTimeout:  o.HTTP.ReadTimeout,
		WriteTimeout: o.HTTP.WriteTimeout,
		ShutdownWait: o.HTTP.ShutdownWait,
		Mux:          mux.New(),
	}

	log.Running(o.HTTP.Addr)

	select {
	case err = <-s.Start():
		log.Error(err)
		log.Stopped()
	case sig := <-signal.Wait():
		log.GotStopSignal(sig)
		log.Stopped()
	}

	if err = s.Stop(); err != nil {
		log.Error(errors.WithStack(err))
	}

	return errors.WithStack(err)
}
