package tui

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/David-Rushton/console"
)

type app struct {
	Log                  *tuiLogger
	IsRunning            bool
	Rows                 int
	Columns              int
	HasScreenBeenResized bool
}

func NewApp() *app {
	var app = &app{
		Log:       newLogger(),
		IsRunning: true,
	}

	app.Log.Information("Starting up...")
	app.start()

	return app
}

func (a *app) start() {
	go a.startKeyEventListener()
	go a.startConsoleMonitor()
	go a.startSignalEventListener()
}

func (a *app) startKeyEventListener() {
	for a.IsRunning {

	}
}

func (a *app) startConsoleMonitor() {
	var rows, columns int
	var err error

	for a.IsRunning {
		rows, columns, err = console.GetSize()
		if err != nil {
			a.Log.Errorf("Cannot get console screen size because %v\n", err)
		}

		if a.Rows != rows || a.Columns != columns {
			a.Rows = rows
			a.Columns = columns
			a.HasScreenBeenResized = true

			a.Log.Informationf("Console screen resized to %v by %v", a.Columns, a.Rows)
		}
	}
}

func (a *app) startSignalEventListener() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	var exitSignal = <-signals
	a.Log.Informationf("Received singal: %v\n", exitSignal)
	a.Log.Information("Shutting down\n")

	a.IsRunning = false
}
