package main

import "fmt"

//结构体

type person struct {
	name string
	age int
	gender string
	hobby []string
}

func main()  {
	//声明一个person类型的变量p
	var cheng person
	//通过字段赋值
	cheng.name = "cc"
	cheng.age = 25
	cheng.gender = "男"
	cheng.hobby = []string{"双色球","足球"}
	fmt.Println(cheng)
	//访问变量cheng的字段
	fmt.Printf("%T\n", cheng)
	fmt.Println(cheng.name)

	var p2 person
	p2.name = "理想"
	p2.age = 12
	fmt.Printf("type:%T value:%v\n", p2, p2)

	//匿名结构体
	var s struct{
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}