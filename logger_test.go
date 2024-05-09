package logger

import (
	"os"
	"testing"
)

func TestMessage(t *testing.T) {
	logger := &Logger{OutStream: os.Stdout, LowestSeverity: ERROR}
	logger.Error("%s", "Error Simple string")
	logger.Info("%s", "Info Simple string")
}
