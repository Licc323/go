package main

import "fmt"

//变量的作用域
 	var x = 100 //定义一个全局变量
 	//定义一个函数
func f01()  {
	name := "lixiang"
	//函数中查找变量的顺序
	//1. 先在函数内部查找
	//2. 找不到就往函数外面查找，一直找到全局变量
	fmt.Println(x)
	fmt.Println(x, name)
}
func main()  {
	f01()
	//fmt.Println(name)//函数内部定义的变量只能在改函数内部使用
	// 语句块作用域
	//if i := 10; i < 18 {
	//	fmt.Println("haohaoxuexi")
	//}
	//fmt.Println(i)
	//for j := 0; j < 5; j++ {
	//	fmt.Println(j)
	//}
	//fmt.Println(j)
}