package main

import "fmt"

//递归
//为了理解递归是什么，你首先要搞懂什么是递归
//递归适合处理那种问题相同；问题的规模越来越小的场景
//递归一定要给出一个明确的退出条件

func f(n uint64) uint64{
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

//上台阶，n个台阶，一次可以走一步，也可以走两步有多少种
func taijie(n uint64) uint64 {
	if n == 1 {
		//如果只有一个台阶只有一种走法
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}
func main()  {
	ret := f(5)
    fmt.Println(ret)
	ret2 := taijie(4)
	fmt.Println(ret2)
}
