package main

import "fmt"

//匿名函数
//   var l1 = func(x, y int)  {
//	fmt.Println(x + y)
//}
func main() {
//函数内部没有办法声明带名字的函数
	l1 := func(x, y int) {
		fmt.Println(x + y)
	}
	l1(10, 20)
	//如果只是调用一次的函数，还可以简写为立即执行的函数
	func(x, y int) {
		fmt.Println(x + y)
		fmt.Println("hello world")
	}(100, 200)
}
