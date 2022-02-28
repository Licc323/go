package main

import "fmt"

//map
func main()  {
	var m1 map[string]int
	fmt.Println(m1 == nil)//还没有在内存里面开辟一个空间
	m1 = make(map[string]int, 10)//要估算好map容量避免在程序运行期间在动态扩容
	m1["理想"] = 18
	m1["licc"] = 35

	fmt.Println(m1)
	fmt.Println(m1["理想"])
	//约定成俗用OK返回布尔值
	value, ok := m1["娜扎"] //如果不存在这个key拿到对应值类型的0值
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}
	//map  的遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	//只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "lixiang")
	fmt.Println(m1)
	delete(m1, "沙河")// 删除不存在的

}
