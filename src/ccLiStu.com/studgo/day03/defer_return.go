package main

import "fmt"

//go中的函数的return不是原子操作，在底层是分两步来执行的
//1.返回值赋值
//2.真正的ret返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1()int  {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回值 = x
}
func f3() (y int)  {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int)  {
	defer func(x int) {
		x++// 改变的是函数中的副本
	}(x)
	return 5// 返回值x = 5
}
func main()  {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}