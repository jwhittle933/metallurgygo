package main

import (
	"log"
	"os"

	"github.com/logrusorgru/aurora"
)

var (
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
)

// Log ...
type Log struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

// StartLog ...
func StartLog() *Log {
	trace = log.New(os.Stdout, aurora.Sprintf(aurora.BrightMagenta("[%s] "), "Trace"), log.Ldate|log.Ltime)
	info = log.New(os.Stdout, aurora.Sprintf(aurora.BrightGreen("[%s] "), "Info"), log.Ldate|log.Ltime)
	warning = log.New(os.Stdout, aurora.Sprintf(aurora.BrightYellow("[%s] "), "Warning"), log.Ldate|log.Ltime)
	err = log.New(os.Stdout, aurora.Sprintf(aurora.BrightRed("[%s] "), "Error"), log.Ldate|log.Ltime)

	return &Log{
		Trace:   trace,
		Info:    info,
		Warning: warning,
		Error:   err,
	}
}

// T Trace method
func (l *Log) T(f string, m ...interface{}) {
	l.Trace.Printf(f, m...)
}

// I Info method
func (l *Log) I(f string, m ...interface{}) {
	l.Info.Printf(f, m...)
}

// W Warning method
func (l *Log) W(f string, m ...interface{}) {
	l.Warning.Printf(f, m...)
}

// E Error method
func (l *Log) E(f string, m ...interface{}) {
	l.Error.Printf(f, m...)
}
