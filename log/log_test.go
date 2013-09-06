package log

import (
	"testing"
)

func Test_log(t *testing.T) {
	Trace("hello", "world")
	Debug("hello")
	Info(" hello")
	Warn(" hello")
	Error("hello")
	Fatal("hello")
}
