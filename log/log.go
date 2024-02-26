package log

type Logger struct{}

func New() *Logger {
	return &Logger{}
}

func (lg *Logger) Info(err error)    {}
func (lg *Logger) Fatal(err error)   {}
func (lg *Logger) Warning(err error) {}
func (lg *Logger) Error(err error)   {}
