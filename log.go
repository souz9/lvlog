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
	error_ = log.New(output, "<3>", flags)
	warn   = log.New(output, "<4>", flags)
	info   = log.New(output, "<5>", flags)
	debug  = log.New(output, "<7>", flags)
)

func Erroro(calldepth int, s string) error { return error_.Output(calldepth+1, s) }
func Warno(calldepth int, s string) error  { return warn.Output(calldepth+1, s) }
func Infoo(calldepth int, s string) error  { return info.Output(calldepth+1, s) }
func Debugo(calldepth int, s string) error { return debug.Output(calldepth+1, s) }

func Fatal(v ...interface{}) {
	Erroro(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	Erroro(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	Erroro(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	Erroro(2, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	Warno(2, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	Warno(2, fmt.Sprintf(format, v...))
}

func Info(v ...interface{}) {
	Infoo(2, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	Infoo(2, fmt.Sprintf(format, v...))
}

func Debug(v ...interface{}) {
	Debugo(2, fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	Debugo(2, fmt.Sprintf(format, v...))
}
