package main

import (
	"fmt"
	"sync"
	"time"
)

//为什么需要context？
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)

func f1()  {
	defer wg.Done()
FORLOOP:
	for  {
			fmt.Println("诗妹")
			time.Sleep(time.Millisecond * 500)
			select {
			case <-exitChan:
				break FORLOOP
			default:
			}
		}
}

func main()  {

	wg.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	//如何通知子groutine退出
	exitChan <- true
	wg.Wait()
}