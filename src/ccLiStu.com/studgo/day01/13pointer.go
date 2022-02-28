package main

import "fmt"

//vscode 不支持go mod

//pointer

func main()  {
	//1. &：取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	//2. *：根据地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	//new函数深情一个内存地址
	var a1 *int //nil pointer
	fmt.Println(a1)
	var a2 = new(int)  //new 函数申请一个内存地址
	//接下来 我不争不抢，留在我身边的人都是心甘情愿的。
	fmt.Println(a2)
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)
	//make 和 new
	//1.make和new都是用来申请内存的
	//new很少用，一般用来给基本数据类型申请内存，string、int返回的是对应类型的指针
	//make是用来给slice、map、channel申请内存的，make函数返回的是对应的这三个类型



}