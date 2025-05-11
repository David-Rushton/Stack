package tui

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/David-Rushton/console"
	"golang.org/x/term"
)

type app struct {
	Log                  *tuiLogger
	IsRunning            bool
	Rows                 int
	Columns              int
	HasScreenBeenResized bool

	consoleStartingState *term.State
}

func NewApp() *app {
	var app = &app{
		Log:       newLogger(),
		IsRunning: true,
	}

	app.start()

	return app
}

func (a *app) start() {
	a.Log.Information("Starting up...")

	console.HideCursor()

	// Enter raw mode
	var fd = int(os.Stdin.Fd())
	var err error
	a.consoleStartingState, err = term.MakeRaw(fd)
	if err != nil {
		a.Log.Errorf("Cannot enter raw mode because %v\n", err)
	}

	go a.startKeyEventListener()
	go a.startConsoleMonitor()
	go a.startSignalEventListener()
}

func (a *app) stop() {
	a.Log.Information("Shutting down\n")

	a.IsRunning = false

	// Exit raw mode
	if a.consoleStartingState != nil {
		if err := term.Restore(int(os.Stdin.Fd()), a.consoleStartingState); err != nil {
			a.Log.Errorf("Cannot exit raw mode because %v\n", err)
		}
	}

	console.ShowCursor()
}

func (a *app) startKeyEventListener() {
	for a.IsRunning {
		readKey(a)
	}
}

func (a *app) startConsoleMonitor() {
	var rows, columns int
	var err error

	for a.IsRunning {
		rows, columns, err = console.GetSize()
		if err != nil {
			// TODO: Could lead to log spam.  At some point should we fail?
			a.Log.Errorf("Cannot get console screen size because %v\n", err)
		}

		if a.Rows != rows || a.Columns != columns {
			a.Rows = rows
			a.Columns = columns
			a.HasScreenBeenResized = true

			a.Log.Informationf("Console screen resized to %v by %v", a.Columns, a.Rows)
		}

		time.Sleep(time.Millisecond * 350)
	}
}

func (a *app) startSignalEventListener() {
	defer a.stop()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	var exitSignal = <-signals
	a.Log.Informationf("Received singal: %v\n", exitSignal)
	a.stop()
}
