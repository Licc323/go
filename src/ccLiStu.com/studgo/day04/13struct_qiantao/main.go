package main

import "fmt"

//结构体嵌套
type address struct {
	province string
	city string
}
type workPlace struct {
	province string
	city string
}
type person struct {
	name string
	age int
	//addr address
	address //匿名嵌套结构体
	workPlace
}
type company struct {
	name string
	addr address

}

func main()  {
	p1 := person{
		name: "chengcheng",
		age: 18,
		address: address{
			province: "浙江",
			city: "杭州",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.address.city)
	//fmt.Println(p1.city)//先在自己结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	fmt.Println(p1.workPlace.city)
	fmt.Println(p1.address.city)
}