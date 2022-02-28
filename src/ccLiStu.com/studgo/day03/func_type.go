package main

import "fmt"

//函数类型
func c1()  {
	fmt.Println("hello 沙河")
}
func c2() int {
	return 10
}
//函数也可以作为参数类型
func c3(x func()int)  {
	ret := x()
	fmt.Println(ret)
}
func c4(x, y int) int {
	return x + y
}
//函数还可以作为返回值
func cc(a, b int) int {
	return a + b
}
func c5(x func()int)func(int, int) int {
	return cc
}

func main()  {
	a := c1
	fmt.Printf("%T\n", a)
	b := c2
	fmt.Printf("%T\n", b)

	c3(c2)
	c3(b)
	fmt.Printf("%T\n", c4)
	//c3(c4)
	c7 :=c5(c2)
	fmt.Printf("%T\n", c7)
}
//函数的定义：1.基本格式
//参数格式：有参数的函数、有参数类型简写、可变参数
//返回值的格式：有返回值、多返回值、命名返回值
//高阶函数：函数也是一种类型，他可以作为参数也可以作为返回值
//变量作用域：1.全局作用域、2.函数作用域：先在函数内部找变量，找不到往外层找。函数内部的变量，外部是访问不到的.
//3.代码块作用域
//匿名函数：没有名字的函数