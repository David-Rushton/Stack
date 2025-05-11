package tui

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

type tuiLogger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
}

// Creates a logger than output each log entry to:
//   - A file located in the users home director
//   - In-memory repository
func newLogger() *tuiLogger {
	// TODO: Path should be configurable.
	var homeDir, err = os.UserHomeDir()
	if err != nil {
		// TODO: There should be a fallback here, rather than a hard fail.
		log.Fatalf("Cannot initialise logger because %v\n", err)
	}
	var logFileName = fmt.Sprintf("stack.%v.log", time.Now().Format("20060102-1504"))
	var logDirectory = path.Join(homeDir, ".stack")
	err = os.Mkdir(logDirectory, 0755)
	if err != nil {
		if !os.IsExist(err) {
			// TODO: There should be a fallback here, rather than a hard fail.
			log.Fatalf("Cannot initialise logger because %v %T\n", err, err)
		}
	}
	var logPath = path.Join(homeDir, ".stack", logFileName)
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		// TODO: There should be a fallback here, rather than a hard fail.
		log.Fatalf("Cannot initialise logger because %v\n", err)
	}

	var inMemoryLog = InMemoryLog{}

	var multiWriter = io.MultiWriter(logFile, &inMemoryLog)

	return &tuiLogger{
		infoLogger:    log.New(multiWriter, "info   > ", log.Default().Flags()),
		warningLogger: log.New(multiWriter, "WARN   > ", log.Default().Flags()),
		errorLogger:   log.New(multiWriter, "ERROR  > ", log.Default().Flags()),
		fatalLogger:   log.New(multiWriter, "FATAL  > ", log.Default().Flags()),
		panicLogger:   log.New(multiWriter, "PANIC! > ", log.Default().Flags()),
	}
}

func (l *tuiLogger) Information(v ...any) {
	l.infoLogger.Print(v...)
}

func (l *tuiLogger) Informationf(format string, v ...any) {
	l.infoLogger.Printf(format, v...)
}

func (l *tuiLogger) Informationln(v ...any) {
	l.infoLogger.Println(v...)
}

func (l *tuiLogger) Warning(v ...any) {
	l.warningLogger.Print(v...)
}

func (l *tuiLogger) Warningf(format string, v ...any) {
	l.warningLogger.Printf(format, v...)
}

func (l *tuiLogger) Warningln(v ...any) {
	l.warningLogger.Println(v...)
}

func (l *tuiLogger) Error(v ...any) {
	l.errorLogger.Print(v...)
}

func (l *tuiLogger) Errorf(format string, v ...any) {
	l.errorLogger.Printf(format, v...)
}

func (l *tuiLogger) Errorln(v ...any) {
	l.errorLogger.Println(v...)
}

func (l *tuiLogger) Fatal(v ...any) {
	l.fatalLogger.Fatal(v...)
}

func (l *tuiLogger) Fatalf(format string, v ...any) {
	l.fatalLogger.Fatalf(format, v...)
}

func (l *tuiLogger) Fatalln(v ...any) {
	l.fatalLogger.Fatalln(v...)
}

func (l *tuiLogger) Panic(v ...any) {
	l.panicLogger.Panic(v...)
}

func (l *tuiLogger) Panicf(format string, v ...any) {
	l.panicLogger.Panicf(format, v...)
}

func (l *tuiLogger) Panicln(v ...any) {
	l.panicLogger.Panicln(v...)
}
