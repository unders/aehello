package http

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Server struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	ShutdownWait time.Duration
	Mux          http.Handler

	srv *http.Server
}

func (s Server) Start() chan error {
	s.srv = &http.Server{
		Addr:         s.Addr,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
		Handler:      s.Mux,
	}

	ch := make(chan error, 1)

	go func() {
		// ListenAndServe always returns a non-nil error.
		ch <- s.srv.ListenAndServe()
	}()

	return ch
}

func (s Server) Stop() error {
	if s.srv == nil {
		return nil
	}

	ctx, _ := context.WithTimeout(context.Background(), s.ShutdownWait)
	err := s.srv.Shutdown(ctx)

	return errors.WithStack(err)
}
