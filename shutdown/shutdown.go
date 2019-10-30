package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

type Configuration struct {
	BeforeExit func(string)
	Signals    []os.Signal
}

var defaultSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
var exitChan = make(chan struct{ message string })

func WaitExit(config *Configuration) {

	sigChan := make(chan os.Signal, 1)

	if config != nil {
		if len(config.Signals) > 0 {
			defaultSignals = config.Signals
		}
	}

	signal.Notify(sigChan, defaultSignals...)

	select {
	case s := <-exitChan:
		onExit(s.message, config)
	case s := <-sigChan:
		onExit(s.String(), config)
	}
}

func onExit(s string, config *Configuration) {

	defer func() {
		_ = recover()
	}()

	if config != nil && config.BeforeExit != nil {
		config.BeforeExit(s)
	}
}

func Exit(msg string) {
	exitChan <- struct{ message string }{msg}
}
