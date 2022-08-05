package main

import (
	"log"
	"os"
	"time"

	"github.com/msrexe/grace"
)

// Example of timeout reached and force shutdown.
func main() {
	// Shutdown the server with a timeout of 5 seconds.
	go func() {
		grace.ShutdownWithTimeout(5*time.Second, func() {
			time.Sleep(10 * time.Second)
			log.Println("Application gracefully stopped")
		}, os.Interrupt)
	}()

	for {
		log.Println("Hello World")
		time.Sleep(2 * time.Second)
	}
}
