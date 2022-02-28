package main

import (
	"flag"
	"fmt"
)

//flag 获取命令行参数

func main()  {
	//创建一个标志位参数
	name := flag.String("name", "雨诗", "请输入名字")
	age := flag.Int("age", 23, "请输入真实的年龄")
	married := flag.Bool("married", false, "结婚了吗")
	//使用flag
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)

	fmt.Println(flag.Args())  // 返回命令行参数后的其他参数
	fmt.Println(flag.NArg())  // 返回命令行参数后的其他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数

}