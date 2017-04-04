package log

import "fmt"

// log messages
const (
	started       = "\n** STARTED **\ninfo=started app=helloworld\n"
	release       = "info=version:%s buildstamp:%s githash:%s app=helloworld\n"
	running       = "info=listens on address %s app=helloworld\n"
	hostname      = "info=hostname:%s app=helloworld\n"
	stopping      = "info=stopping app=helloworld"
	stopped       = "info=stopped app=helloworld\n** STOPPED **\n"
	gotStopSignal = "\ninfo=got signal %s app=helloworld\n"
	errFormat     = "err=%s\nstacktrace=%+v app=hellworld\n"
)

// Started logs started info
func Started() {
	writeInfo(started)
}

// Release logs release info
func Release(version, buildStamp, gitHash string) {
	msg := fmt.Sprintf(release, version, buildStamp, gitHash)
	writeInfo(msg)
}

// Running logs running info
func Running(addr string) {
	msg := fmt.Sprintf(running, addr)
	writeInfo(msg)
}

// Host logs running info
func Host(host string) {
	msg := fmt.Sprintf(hostname, host)
	writeInfo(msg)
}

// GotStopSignal logs got stop signal info
func GotStopSignal(signal fmt.Stringer) {
	msg := fmt.Sprintf(gotStopSignal, signal)
	writeInfo(msg)
}

// Config logs config info
func Config(config string) {
	writeInfo(config)
}

// Stopped logs stop info
func Stopped() {
	writeInfo(stopped)
}

// Stopping logs stopping info
func Stopping() {
	writeInfo(stopping)
}

// Info logs msg to info
func Info(info string) {
	writeInfo(info)
}

// Error logs errors
func Error(err error) {
	msg := fmt.Sprintf(errFormat, err, err)
	writeError(msg)
}
