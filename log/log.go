package log

import (
	"fmt"
	"os"
)

const (
	RED   = "\033[31m"
	RESET = "\033[0m"
)

type Logger struct{}

// New used for generate new logger instance of Logger struct
func New() *Logger {
	return &Logger{}
}

// Info will show information messages in stdout
func (lg *Logger) Info(err error) {}

// Fatal will show an error and exit with not success
// status code
func (lg *Logger) Fatal(err error) {}

// Warning
func (lg *Logger) Warning(err error) {
	fmt.Fprintf(os.Stderr, RED+"%s\n"+RESET, err.Error())
}

// Error
func (lg *Logger) Error(err error) {}
