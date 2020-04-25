package goui_tools

import "fmt"

type Logger interface {
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
}

type internal struct {
}

func (i internal) Info(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (i internal) Warning(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

func (i internal) Error(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

var logger Logger

func SetLogger(l Logger) {
	logger = l
}

func getLogger() Logger {
	if logger == nil {
		logger = new(internal)
	}
	return logger
}
