package signalwatcher

import (
	"os"
	"os/signal"
	"syscall"
)

func Watch(callback func(Signal)) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		// This will block until a signal is received
		<-signals

		callback(QUIT)
		Watch(callback)
	}()
}