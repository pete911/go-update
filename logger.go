package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	calldepth   = 2
	prefix      = "[INFO]  "
	errPrefix   = "[ERROR] "
	fatalPrefix = "[FATAL] "
	flags       = log.Ltime
)

var (
	InfoLogger  = newStdOutLogger(os.Stdout, prefix)
	ErrorLogger = newStdErrLogger(os.Stderr, errPrefix)
	FatalLogger = newStdErrLogger(os.Stderr, fatalPrefix)
)

func newStdOutLogger(out io.Writer, prefix string) *log.Logger {
	return log.New(out, prefix, flags)
}

func newStdErrLogger(out io.Writer, prefix string) *log.Logger {
	return log.New(out, prefix, flags)
}

func Fatal(v ...interface{}) {
	FatalLogger.Output(calldepth, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	FatalLogger.Output(calldepth, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Error(v ...interface{}) {
	ErrorLogger.Output(calldepth, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	ErrorLogger.Output(calldepth, fmt.Sprintf(format, v...))
}

func Log(v ...interface{}) {
	InfoLogger.Output(calldepth, fmt.Sprint(v...))
}

func Logf(format string, v ...interface{}) {
	InfoLogger.Output(calldepth, fmt.Sprintf(format, v...))
}
