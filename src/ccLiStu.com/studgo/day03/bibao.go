package main

import "fmt"

//闭包
func d1(d func())  {
	fmt.Println("this is d1")
	d()
}
func d2(x, y int)  {
	fmt.Println("this is d2")
	fmt.Println(x + y)
}
//目的 通过d1调用d2的参数
//定义一个函数对d2进行包装
func d3(d func(int, int), x, y int) func() {
	tmp := func() {
		fmt.Println(x + y)
		//f2()
		d(x, y)
	}
	return tmp
}

func main() {
	ret := d3(d2,100, 200)//把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	d1(ret)
}
