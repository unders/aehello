package config

import (
	"flag"
	"os"
	"time"

	"github.com/pkg/errors"
)

var localMachine = false

// Options contains the configuration
type Options struct {
	Name          string
	Env           env
	HTTP          http
	GCloudProject string
	GAE           gae
}

type gae struct {
	Service  string
	Version  string
	Instance string
}

// HTTP contains the http configuration
type http struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	ShutdownWait time.Duration
}

type env string

const local = "local"

func (e env) IsLocal() bool {
	return string(e) == local
}

// New returns config.Options
func New(args []string) (Options, error) {
	o := Options{Name: "Helloworld"}
	setFromArgs(args, &o)
	err := setFromEnv(&o)

	return o, errors.WithStack(err)
}

func setFromArgs(args []string, o *Options) {
	addr := "0.0.0.0:8080"
	l := false
	flag.StringVar(&addr, "http", addr, "HTTP service address.")
	flag.BoolVar(&l, "l", l, "Running on local machine")
	flag.Parse()

	localMachine = l
	o.HTTP = http{
		Addr:         addr,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		ShutdownWait: 5 * time.Second,
	}
}

func setFromEnv(o *Options) error {
	// GOOGLE_APPLICATION_CREDENTIALS
	if localMachine {
		o.Env = local
		return nil
	}

	p := os.Getenv("GCLOUD_PROJECT")
	if p == "" {
		return errors.New("ENV: GCLOUD_PROJECT is not set")
	}
	o.GCloudProject = p

	i := os.Getenv("GAE_INSTANCE")
	if i == "" {
		return errors.New("ENV: GAE_INSTANCE is not set")
	}
	o.GAE.Instance = i

	s := os.Getenv("GAE_SERVICE")
	if s == "" {
		return errors.New("ENV: GAE_SERVICE is not set")
	}
	o.GAE.Service = s

	v := os.Getenv("GAE_VERSION")
	if v == "" {
		return errors.New("ENV: GAE_VERSION is not set")
	}
	o.GAE.Version = v

	e := os.Getenv("ENVIRONMENT")
	if v == "" {
		return errors.New("ENV: ENVIRONMENT is not set")
	}
	o.Env = env(e)

	return nil
}
