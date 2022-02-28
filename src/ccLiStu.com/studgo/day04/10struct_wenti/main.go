package main

import "fmt"

//结构体遇到的问题
//1.myInt(100)是什么？
type myTint int

func (m myTint) hello() {
	fmt.Println("我是一个int")
}

func main()  {
	//声明一个int32类型的变量x，他的值是10
	var x int32 = 10
	//x := int32(10)
	fmt.Println(x)
	//声明一个myInt类型的变量m，他的值是100
	var m = myTint(100)
	fmt.Println(m)

	//结构体初始化
	type person struct {
		name string
		age int
	}
	//方法1 ：
	var p person
	p.name = "chengcheng"
	p.age = 19
	fmt.Println(p)
	var p1 person
	p1.name = "wenwen"
	p1.age = 23
	fmt.Println(p1)
	//方法2：键值对初始化
	//声明同时初始化
	s1 := []int{1, 2, 3, 4}
	m1 := map[string]int{
		"stu1": 100,
		"stu2": 99,
		"stu3": 0,
	}
	fmt.Println(s1, m1)
	var p2 = person{
		name: "xingxing",
		age: 3,
	}
	fmt.Println(p2)
	//值列表初始化
	var p3 = person{
		"lixaing",
		100,
	}
	fmt.Println(p3)
}



//为什么要有构造函数
//func newPerson(name string, age int) person {
//	//别人调用我，我能给他一个person类型的变量
//	return person{
//		name,
//		age,
//	}
//}
