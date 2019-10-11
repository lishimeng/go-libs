package shutdown

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Configuration struct {
	BeforeExit       func(string)
	AuthShutdownTime time.Duration
	Signals          []os.Signal
}

var defaultSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

func WaitExit(config *Configuration) {

	sigChan := make(chan os.Signal, 1)
	exitChan := make(chan struct{ message string })
	if config != nil {
		if len(config.Signals) > 0 {
			defaultSignals = config.Signals
		}
	}

	signal.Notify(sigChan, defaultSignals...)

	go func() {
		if config != nil && config.AuthShutdownTime > time.Second {
			time.Sleep(config.AuthShutdownTime)
			exitChan <- struct{ message string }{"auto"}
		}
	}()
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
