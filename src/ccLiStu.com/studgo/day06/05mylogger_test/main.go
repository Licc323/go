package main

import (
	"ccLiStu.com/studgo/day06/mylogger"
	"time"
)
//测试我们自己写的日志库

var log mylogger.Logger //声明一个全局的接口变量

func main()  {
	log = mylogger.NewConsoleLogger("Info")
	log = mylogger.NewFileLogger("Info", "./", "chengcheng.log", 10*1024*1024)

	for  {
		log.Debug("这是一条Debug日志")
		log.Trace("这是一个trace日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志, id:%d, name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second)
	}
}
