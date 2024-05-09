package logger

import (
	"fmt"
	"io"
)

type Severity string

const (
	DEBUG Severity = "DEBUG"
	INFO  Severity = "INFO"
	ERROR Severity = "ERROR"
	FAIL  Severity = "FAIL"
)

func (s Severity) ToLevel() int {
	switch s {
	case DEBUG:
		return 0
	case INFO:
		return 1
	case ERROR:
		return 2
	default:
		return 3
	}
}

// Logger for Go
type Logger struct {
	Name           string
	OutStream      io.Writer
	LowestSeverity Severity
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.Message(DEBUG, format, args)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.Message(ERROR, format, args)
}

func (l *Logger) Fail(format string, args ...interface{}) {
	l.Message(FAIL, format, args)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.Message(INFO, format, args)
}

func (l *Logger) Message(severity Severity, format string, args ...interface{}) {
	if severity.ToLevel() < l.LowestSeverity.ToLevel() {
		return
	}
	s := fmt.Sprintf(string(severity)+" "+format+"\n", args...)
	fmt.Fprintf(l.OutStream, s)
	if severity == FAIL {
		panic(s)
	}
}
