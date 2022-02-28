package main

import "fmt"

//匿名字段
//字段比较少也比较简洁的场景
type person struct {
	string
	int

}

func main()  {
	p1 := person{
		"chengcheng",
		2000,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}