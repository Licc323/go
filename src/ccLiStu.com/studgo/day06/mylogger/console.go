package mylogger

import (
	"fmt"
	"time"
)

//往终端写日志相关内容
//Logger 日志结构体
//type LogLevel uint16
//
//const (
//	//定义日志级别
//	UNKNOWN LogLevel = iota
//	DEBUG
//	TRACE
//	INFO
//	WARNING
//	ERROR
//	FATAL
//)
//
//func parseLogLevel(s string) (LogLevel, error) {
//	s = strings.ToLower(s)
//	switch s {
//	case "debug":
//		return DEBUG, nil
//	case "trace":
//		return TRACE, nil
//	case "info":
//		return INFO, nil
//	case "wraning":
//		return WARNING, nil
//	case "error":
//		return ERROR, nil
//	case "fatal":
//		return FATAL, nil
//	default:
//		err := errors.New("无效的日志级别")
//		return UNKNOWN, err
//	}
//logger 日志结构体
type  ConsoleLogger struct {
	Level LogLevel
}
//NewLo
//g 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)

	}
}
func (c ConsoleLogger) Debug (fromat string, a ...interface{}) {
		c.log(DEBUG, fromat, a...)
}
func (c ConsoleLogger) Trace (fromat string, a ...interface{}) {
		c.log(TRACE, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg
}

func (c ConsoleLogger) Info (fromat string, a ...interface{}) {
		c.log(INFO, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg
}

func (c ConsoleLogger) Warning (fromat string, a ...interface{}) {
	c.log(WARNING, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [WARNING] %s\n", now.Format("2006-01-02 15:04:05"), msg
}

func (c ConsoleLogger) Error (fromat string, a ...interface{}) {
	c.log(ERROR, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg
}

func (c ConsoleLogger) Fatal (fromat string, a ...interface{}) {
	c.log(FATAL, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg
}
