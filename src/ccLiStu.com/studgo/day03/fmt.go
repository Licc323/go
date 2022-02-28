package main

import "fmt"

//fmt
func main()  {
	//fmt.Print("沙河")
	//fmt.Println("沙河")

	//Printf("格式化字符串"，值)
	//%T:查看类型
	//%d：十进制数
	//%b:二进制数
	//%o:八进制数
	//%x:十六进制数
	//%c:字符
	//%s:字符串
	//%p:指针
	//%f:浮点
	//%t:布尔值
//	var m1 = make(map[string]int, 1)
//	m1["理想"] = 100
//	fmt.Printf("%v\n", m1)
//	fmt.Printf("%#v\n", m1)
//
	printBaifenbi(90)

	//获取用户输入
	var s string
	fmt.Scan(&s)
	fmt.Println("用户输入的内容是：", s)

	var(
		name string
		age int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)

	fmt.Scanln(&name, &age)
	fmt.Println(name, age)
}
func printBaifenbi(num int)  {
	fmt.Printf("%d%%\n",num)
}


