package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/vvidovic/test-go/view"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	setupCloseHandler()

	eventChan := make(chan string, 1)
	view.Init(eventChan)
}

// SetupCloseHandler creates a 'listener' on a new goroutine which will notify the
// program if it receives an interrupt from the OS - so we can cleanup/close termbox.
func setupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	signal.Notify(c, os.Interrupt, syscall.SIGKILL)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP)
	go func() {
		s := <-c
		fmt.Printf("Signal received: %v", s)
		termbox.Close()
		os.Exit(0)
	}()
}
