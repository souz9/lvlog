package log

import (
	"fmt"
	"log"
	"os"
)

var (
	output = os.Stderr
	flags  = log.Lshortfile

	// The Loggers are prefixed with syslog level.
	error = log.New(output, "<3>", flags)
	warn  = log.New(output, "<4>", flags)
	info  = log.New(output, "<5>", flags)
	debug = log.New(output, "<7>", flags)
)

func Fatal(v ...interface{}) {
	error.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	error.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	error.Output(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	error.Output(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	warn.Output(2, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	warn.Output(2, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	info.Output(2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	info.Output(2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	debug.Output(2, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	debug.Output(2, fmt.Sprintf(format, v...))
}
