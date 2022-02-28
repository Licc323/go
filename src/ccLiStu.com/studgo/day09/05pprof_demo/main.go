package main

import "fmt"

//一段有问题的代码

func logicCode()  {
	var c chan int
	for  {
		select {  //多路复用
		case v := <-c:
			fmt.Printf("recv from chan, value:%v\n", v)
		default:
		}
	}
}
