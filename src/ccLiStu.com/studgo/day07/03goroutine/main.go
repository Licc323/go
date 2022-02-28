package main

import (
	"fmt"
)

//goroutine

func hello(i int)  {
	fmt.Println("hello", i)
}

//程序启动之后 去创建一个 主goroutine去执行
func main()  {
	for i := 0; i<1000; i++ {
		//go hello(i) //开启一个单独的goroutine 去执行hello函数
		go func(i int) {
			fmt.Println(i)//  用的是函数参数的那个i，不是外面的i
		}(i)
	}
	fmt.Println("main")
	//time.Sleep(time.Second)
	//main函数结束了 由main函数启动的goroutine 也都结束了
}
