package config

import (
	"flag"
	"os"

	"github.com/pkg/errors"
)

var localMachine = false

// Options contains the configuration
type Options struct {
	Name          string
	Env           env
	Server        server
	GCloudProject string
	GAE           gae
}

type gae struct {
	service  string
	version  string
	instance string
}

// Server contains the server configuration
type server struct {
	Addr string
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
	o.Server.Addr = addr
}

func setFromEnv(o *Options) error {
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
	o.GAE.instance = i

	s := os.Getenv("GAE_SERVICE")
	if s == "" {
		return errors.New("ENV: GAE_SERVICE is not set")
	}
	o.GAE.service = s

	v := os.Getenv("GAE_VERSION")
	if v == "" {
		return errors.New("ENV: GAE_VERSION is not set")
	}
	o.GAE.version = v

	e := os.Getenv("ENVIRONMENT")
	if v == "" {
		return errors.New("ENV: ENVIRONMENT is not set")
	}
	o.Env = env(e)

	return nil
}
