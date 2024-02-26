package log

import (
	"fmt"
	"os"
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
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
}

// Error
func (lg *Logger) Error(err error) {}
