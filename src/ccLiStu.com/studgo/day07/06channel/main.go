package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int  //需要制定同道中元素的类型
var wg sync.WaitGroup

func noBufChannel()  {
	b = make(chan int) //不带缓冲区的
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10   //hang住了
	fmt.Println("10发送到通道b中...")
	wg.Wait()
}

func bufChannel()  {
	fmt.Println(b) //nil
	b = make(chan int, 16) //通道初始化  //16: 带缓冲区的
	b <- 10
	fmt.Println("10发送到通道b中...")
	b <- 20
	fmt.Println("20发送到通道b中...")
	x := <-b
	fmt.Println("从通道b中取到了", x)
	close(b)
}
func main()  {
	bufChannel()
}
