package tui

import (
	"os"
	"os/signal"
	"syscall"
)

type app struct {
	Log       *tuiLogger
	IsRunning bool
}

func NewApp() *app {
	var app = &app{
		Log:       newLogger(),
		IsRunning: true,
	}

	app.Log.LogInformation("Starting up...")
	app.start()

	return app
}

func (a *app) start() {
	go a.startKeyEventListener()
	go a.startSignalEventListener()
	go a.startConsoleMonitor()
}

func (a *app) startKeyEventListener() {
	for a.IsRunning {

	}
}

func (a *app) startSignalEventListener() {
	for a.IsRunning {

	}
}

func (a *app) startConsoleMonitor() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	var exitSignal = <-signals
	a.Log.LogInformationf("Shutting down (%v)\n", exitSignal)

	a.IsRunning = false
}
