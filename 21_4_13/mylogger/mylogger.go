package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//向终端写入日志内容
type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

var levelTostr = map[int]string{
	int(UNKNOWN): "unkonwn",
	int(DEBUG):   "DEBUG",
	int(TRACE):   "TRACE",
	int(INFO):    "INFO",
	int(WARNING): "WARNING",
	int(ERROR):   "ERROR",
	int(FATAL):   "FATAL",
}

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

// Logger
type ConsoleLogger struct {
	Level LogLevel
}

type Logger interface {
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Error(format string, a ...interface{})
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {

	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}

	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return
}
