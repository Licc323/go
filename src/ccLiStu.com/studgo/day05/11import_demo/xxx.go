package main

import (
	"ccLiStu.com/studgo/day05/10package"
	"fmt"
)

var x = 100 //全局变量

const pi = 3.14 //声明：
func init()  {
	fmt.Println("自动执行！")
	fmt.Println(x, pi)
}
func main()  {
	ret := calc.Add(10, 20)
	fmt.Println(ret)
}