package main


import (
	"context"
	"fmt"
	"sync"
	"time"
)

//为什么需要context？
var wg sync.WaitGroup
var exitChan = make(chan bool, 1)
func f2(ctx context.Context)  {
	defer wg.Done()
FORLOOP:
	for  {
		fmt.Println("夏雨诗")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f1(ctx context.Context)  {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for  {
		fmt.Println("诗妹")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main()  {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f1(ctx)
	time.Sleep(time.Second * 5)
	//如何通知子groutine退出
	cancel()
	wg.Wait()
}
//