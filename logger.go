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
	FatalLogger.Output(calldepth, colorFmt(31, fmt.Sprint(v...)))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	FatalLogger.Output(calldepth, colorFmt(31, fmt.Sprintf(format, v...)))
	os.Exit(1)
}

func Error(v ...interface{}) {
	ErrorLogger.Output(calldepth, colorFmt(31, fmt.Sprint(v...)))
}

func Errorf(format string, v ...interface{}) {
	ErrorLogger.Output(calldepth, colorFmt(31, fmt.Sprintf(format, v...)))
}

func Log(v ...interface{}) {
	InfoLogger.Output(calldepth, colorFmt(34, fmt.Sprint(v...)))
}

func Logf(format string, v ...interface{}) {
	InfoLogger.Output(calldepth, colorFmt(34, fmt.Sprintf(format, v...)))
}

func PrintPrompt(v ...interface{}) {
	fmt.Printf(colorFmt(36, fmt.Sprint(v...)))
}

func PrintPromptf(format string, v ...interface{}) {
	fmt.Printf(colorFmt(36, fmt.Sprintf(format, v...)))
}

func Print(v ...interface{}) {
	fmt.Printf(colorFmt(35, fmt.Sprint(v...)))
}

func Printf(format string, v ...interface{}) {
	fmt.Printf(colorFmt(35, fmt.Sprintf(format, v...)))
}

func colorFmt(code int, msg string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", code, msg)
}
