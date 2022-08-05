package grace

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	DefaultSignals = []os.Signal{os.Interrupt, syscall.SIGINT, syscall.SIGTERM}
	DefaultTimeout = time.Minute
)

// Shutdown is a function that will listen for a signal and then
// shutdown the application.
func Shutdown(shutdownFunc func(), signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultSignals
	}

	notifyCh := make(chan os.Signal, 2)
	signal.Notify(notifyCh, signals...)

	log.Printf("received '%s' signal, terminating..", <-notifyCh)

	go func() {
		select {
		case <-time.After(DefaultTimeout):
			log.Fatalf("default timeout reached: terminating..")
		case sig := <-notifyCh:
			log.Fatalf("received second signal %s: terminating..", sig)
		}
	}()

	shutdownFunc()
}

// ShutdownWithTimeout is a function that will listen for a signal and then
// shutdown the application. It will also set a timeout for the shutdown.
func ShutdownWithTimeout(timeout time.Duration, shutdownFunc func(), signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultSignals
	}

	notifyCh := make(chan os.Signal, 2)
	signal.Notify(notifyCh, signals...)

	log.Printf("received '%s' signal", <-notifyCh)
	log.Printf("terminating in %s", timeout)

	go func() {
		select {
		case <-time.After(timeout):
			log.Fatalf("timeout reached: terminating..")
		case sig := <-notifyCh:
			log.Fatalf("received second signal %s: terminating..", sig)
		}
	}()

	shutdownFunc()
}
