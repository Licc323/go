package main

import (
	"fmt"
)

func main()  {
	//指针
	//golang只能读不能修改，不能修改指针变量指向的地址
	addr := "沙河"
	addrP := &addr
	fmt.Println(addrP) // 内存地址
	fmt.Printf("%T\n", addrP) //类型
	addrV := *addrP //根据内存地址找值
	fmt.Println(addrV)

	//map
	//存放指针的键值对
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["lixiang"] = 100
	fmt.Println(m1)
	fmt.Println(m1["ji"]) //如果key不存在返回的是value对应类型的零值
	// 如果返回值是布尔型，我们通常用OK去接收它
	score, ok := m1["ji"]
	if !ok {
		fmt.Println("没有姬无命这个人")
	} else {
		fmt.Println("姬无命的分数是", score)
	}
	delete(m1, "理想")// 删除的key不存在什么都不干
	delete(m1, "lixiang")
	fmt.Println(m1)
	fmt.Println(m1 == nil)// 重新开辟了内存
}
//从此你又是谁的XXX
//自控多情损梵行，入山又怕误倾城，世间安得双全法，不负如来不负卿