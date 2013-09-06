package log

import (
	"fmt"
	"github.com/Bitnick2002/goa/console"
	"time"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

const (
	traceColor = console.BLUE
	debugColor = console.CYAN
	infoColor  = console.GREEN
	warnColor  = console.YELLOW
	errorColor = console.RED
	fatalColor = console.MAGENTA
)

var Level int

const timeFormat = "2006-01-02 15:04:05.000"

type Printer struct {
	PrintTrace func(message string)
	PrintDebug func(message string)
	PrintInfo  func(message string)
	PrintWarn  func(message string)
	PrintError func(message string)
	PrintFatal func(message string)
}

var printer Printer

func init() {
	printer.PrintTrace = func(message string) {
		console.Color(traceColor).Println(message)
	}
	printer.PrintDebug = func(message string) {
		console.Color(debugColor).Println(message)
	}
	printer.PrintInfo = func(message string) {
		console.Color(infoColor).Println(message)
	}
	printer.PrintWarn = func(message string) {
		console.Color(warnColor).Println(message)
	}
	printer.PrintError = func(message string) {
		console.Color(errorColor).Println(message)
	}
	printer.PrintFatal = func(message string) {
		console.Color(fatalColor).Println(message)
	}
}
func Trace(a ...interface{}) {
	if Level <= TRACE {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[TRACE]", fmt.Sprint(a...))
		printer.PrintTrace(message)
	}
}

func Debug(a ...interface{}) {
	if Level <= DEBUG {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[DEBUG]", fmt.Sprint(a...))
		printer.PrintDebug(message)
	}
}
func Info(a ...interface{}) {
	if Level <= INFO {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[INFO]", fmt.Sprint(a...))
		printer.PrintInfo(message)
	}
}
func Warn(a ...interface{}) {
	if Level <= WARN {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[WARN]", fmt.Sprint(a...))
		printer.PrintWarn(message)
	}
}
func Error(a ...interface{}) {
	if Level <= ERROR {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[ERROR]", fmt.Sprint(a...))
		printer.PrintError(message)
	}
}
func Fatal(a ...interface{}) {
	if Level <= FATAL {
		now := time.Now().Format(timeFormat)
		message := fmt.Sprintf("%s %s %s", now, "[FATAL]", fmt.Sprint(a...))
		printer.PrintFatal(message)
	}
}
