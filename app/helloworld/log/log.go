package log

// Documentation
// https://cloud.google.com/appengine/docs/standard/go/logs/
// https://cloud.google.com/logging/docs/reference/libraries

import (
	"context"
	"io"
	"log"

	"cloud.google.com/go/logging"
)

func init() {
	log.SetFlags(0)
}

// Closer closes the client
type Closer struct {
	client *logging.Client
}

// Close writes close error (if any) to stderr
// since we cannot write to the client logger we try to close
func (c Closer) Close() {
	if c.client != nil {
		if err := c.client.Close(); err != nil {
			log.Print("ERROR ", err)
		}
	}
}

// Close logs if there is a close error
func Close(c io.Closer) {
	if err := c.Close(); err != nil {
		Error(err)
	}
}

var out *logging.Logger
var useBackupLog = true

// Init initialize the logging client for the application
func Init(projectID string, localMachine bool) Closer {
	c := Closer{}

	if localMachine {
		return c
	}

	client, err := logging.NewClient(context.Background(), projectID)
	if err != nil {
		log.Println("Could not create logging Client", err)
		return c
	}
	c.client = client

	if err := client.Ping(context.TODO()); err != nil {
		log.Println("ERROR Could not Ping logging Server; error: ", err)
		log.Printf("STACKTRACE  %+v\n", err)
		return c
	}

	client.OnError = func(err error) {
		log.Println("ERROR ", err)
	}

	out = client.Logger("helloworld")

	useBackupLog = false
	return c
}

func writeError(p string) {
	if useBackupLog {
		log.Print("ERROR ", p)
		return
	}

	e := logging.Entry{
		Severity: logging.Error,
		Payload:  p,
	}

	out.Log(e)
}

func writeInfo(p string) {
	if useBackupLog {
		log.Print(p)
		return
	}

	e := logging.Entry{
		Severity: logging.Info,
		Payload:  p,
	}

	out.Log(e)
}
