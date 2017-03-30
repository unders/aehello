package signal

import (
	"os"
	"os/signal"
	"syscall"
)

// Wait returns an os.Signal chan that waits for these signals:
//
//      * SIGTERM   - request for termination (process can release its resources).
//      * SIGINT    - sent from terminal when it wishes to interrupt the process.
//
// Usage:
//
//       sig := <-signal.Wait()
//
func Wait() chan os.Signal {
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt,
		syscall.SIGTERM,
		syscall.SIGINT)

	return interrupt
}
