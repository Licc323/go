package main

import "fmt"

//函数：一段代码的封装

func f12()  {
	fmt.Println("hello 沙河")
}
func f22(name string) {
	fmt.Println("Hello", name)
}
func f32(x int, y int)int  {
	sum := x + y
	return sum
}
//参数类型简写
func f42(x,y int) int  {
	return x + y
}
func f5(title string, y ...int) int  {
	fmt.Println(y) //y是一个int类型的切片
	return 1
}
//命名返回值
func f6(x,y int)(sum int)  {
	sum = x + y //如果使用命名的返回值，那么在函数中可以直接使用返回值变量
	return //如果使用命名返回值 return后面可以省略返回值变量

}
//go 语言中支持多个返回值
func f7(x, y int)(sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main()  {
	f12()
	f22("理想")
	f22("姬无命")
	f32(100,200)//调用函数

	ret := f32(100, 200)
	fmt.Println(ret)

	f5("lixiang", 1,2,3,4,5)
//在一个命名的函数中不允许在命名
}
