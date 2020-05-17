package logging

import (
	"fmt"
	"log"
	"os"
)

// CallDepth for Log
const (
	CallFromOuter = 2
	CallFromInner = 3
)

// Logging level
const (
	NOTSET = 8 << iota
	DEBUG
	INFO
	WARNING
	ERROR
	CRITICAL
)

// getPrefix return the prefix string of given level
func getPrefix(level int) string {
	prefix := ""
	switch level {
	case DEBUG:
		prefix = "DEBUG"
	case INFO:
		prefix = "INFO"
	case WARNING:
		prefix = "WARNING"
	case ERROR:
		prefix = "ERROR"
	case CRITICAL:
		prefix = "CRITICAL"
	default:
		prefix = "NOTSET"
	}
	return fmt.Sprintf("[%s] ", prefix)
}

// Log write a log
func Log(level int, callDepth int, v ...interface{}) {
	l := log.Logger{}
	l.SetOutput(os.Stdout)
	l.SetPrefix(getPrefix(level))
	l.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	_ = l.Output(callDepth, fmt.Sprintln(v...))
}

func Debug(v ...interface{}) {
	Log(DEBUG, CallFromInner, v...)
}

func Info(v ...interface{}) {
	Log(INFO, CallFromInner, v...)
}

func Warning(v ...interface{}) {
	Log(WARNING, CallFromInner, v...)
}

func Error(v ...interface{}) {
	Log(ERROR, CallFromInner, v...)
}

func Critical(v ...interface{}) {
	Log(CRITICAL, CallFromInner, v...)
}

func Notset(v ...interface{}) {
	Log(NOTSET, CallFromInner, v...)
}
