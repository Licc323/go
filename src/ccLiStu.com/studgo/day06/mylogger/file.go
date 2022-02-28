package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里写日志相关代码
//FileLogger 文件日志结构体
type FileLogger struct {
	Level  LogLevel
	filePath string //日志保存的路径
	fileName string // 日志保存的文件名
	fileObj *os.File
	errFileObj *os.File
	maxFileSize int64 
}

//NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	LogLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level: LogLevel,
		filePath: fp,
		fileName: fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile()
	if err != nil {
		panic(nil)
	}
	return fl
}
//根据指定的日志文件路径和文件名打开日志文件
func (f *FileLogger) initFile()(error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}
//判断是否需要记录该日志
func (f FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}
//判断文件是否需要切割
func (f FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get info failed, err:v\n", err)
		return false
	}
	//如果当前文件大小，大于等于日志文件的最大值，就应该返回true
	return fileInfo.Size() > f.maxFileSize
}
//切割文件
func (f *FileLogger) splitFile(file *os.File) (*os.File, error) {
	//需要切割日志文件
	// 备份一下 rename xx.log -> xx.log.bak202111091117
	nowStr := time.Now().Format("20060102150304000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get info failed, err:v\n", err)
		return nil, err
	}

	logName := path.Join(f.filePath, fileInfo.Name())    //拿到当前的日志文件完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)  //拼接一个日志文件备份的名字
	//关闭当前的日志文件
	file.Close()

	os.Rename(logName, newLogName)
	//打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:v\n", err)
		return nil, err
	}

	return fileObj, nil
}
//记录日志的方法
func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			//将打开的新的日志文件对象赋值给newFile
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile, err := f.splitFile(f.errFileObj) //日志文件
				if err != nil {
					return
				}
				//将打开的新的日志文件对象赋值给newFile
				f.errFileObj = newFile
				//如果要记录的日志大于等于error级别，我还要在err日志文件中在记录一遍
				fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
			}
		}
	}
}

func (f FileLogger) Debug (fromat string, a ...interface{}) {
		f.log(DEBUG, fromat, a...)
}
func (f FileLogger) Trace (fromat string, a ...interface{}) {
		f.log(TRACE, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
func (f FileLogger) Info (fromat string, a ...interface{}) {
		f.log(INFO, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [INFO] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
func (f FileLogger) Warning (fromat string, a ...interface{}) {
		f.log(WARNING, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [WARNING] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
func (f FileLogger) Error (fromat string, a ...interface{}) {
		f.log(ERROR, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [ERROR] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}
func (f FileLogger) Fatal (fromat string, a ...interface{}) {
		f.log(FATAL, fromat, a...)
		//now := time.Now()
		//fmt.Printf("[%s] [FATAL] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (f *FileLogger)Close()  {
	f.fileObj.Close()
	f.errFileObj.Close()
}