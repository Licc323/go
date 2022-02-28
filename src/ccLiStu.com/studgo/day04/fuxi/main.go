package main

import "fmt"

func yuanshuai(name string)  {
	fmt.Println("hello", name)
}
func lixiang(f func(string), name string )  {
	f(name)
}
//函数作为返回值
func zhulin() func(int, int)int {
	return func(x, y int)int {
		return x + y
	}
}

func low(f func())  {
	f()
}
//闭包
func bi(f func(string), name string) func() {
	return func() {
		f(name)
	}
}

func main()  {
	lixiang(yuanshuai, "理想")
	ret := zhulin()
	fmt.Printf("%T\n", ret)
	sum :=ret (10, 20)
	fmt.Println(sum)
	//到最后竟庆幸于夕阳西洋仍留在身上
	//来不及讲故事多跌宕

	//闭包实例
	fc := bi(yuanshuai, "元帅")
	low(fc)
}