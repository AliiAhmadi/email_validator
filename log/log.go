package log

import (
	"fmt"
	"os"
)

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (lg *Logger) Info(err error)  {}
func (lg *Logger) Fatal(err error) {}
func (lg *Logger) Warning(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
}
func (lg *Logger) Error(err error) {}
