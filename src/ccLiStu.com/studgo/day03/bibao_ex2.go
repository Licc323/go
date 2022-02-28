package main

import "fmt"

//关键字 函数名（参数）（返回值）{}
func calc(base int) (func(int) int, func(int) int) {
	add1 := func(i int) int {
		base += i
		return base
	}
	sub1 := func(i int) int{
		base -= i
		return base
	}
	return add1, sub1
}
func main()  {
	q1, q2 := calc(10)
	fmt.Println(q1(1), q2(2))//11, 9
	fmt.Println(q1(3), q2(4))//12, 8
	fmt.Println(q1(5), q2(6))//13, 7
}
