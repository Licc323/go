package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//往终端写日志相关内容
//Logger 日志结构体
type LogLevel uint16

//log 接口
type Logger interface {
	Debug (fromat string, a ...interface{})
	Trace (fromat string, a ...interface{})
	Info (fromat string, a ...interface{})
	Warning (fromat string, a ...interface{})
	Error (fromat string, a ...interface{})
	Fatal (fromat string, a ...interface{})
}

const (
	//定义日志级别
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "wraning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case FATAL:
		return "FATAL"
	}
	return "Debug"
}


func getInfo(skip int) (funcName, fileName string, lineNo int){
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)

	return
}

func main()  {

}