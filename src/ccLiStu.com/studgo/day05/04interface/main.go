package main

import "fmt"

//接口实例2
//不管是什么牌子的车都能跑

type car interface {
	run()
}

type falali struct {
	brand string
}

func (f falali)run()  {
	fmt.Printf("%s速度70迈~\n", f.brand)
}
type baoshijie struct {
	brand string
}
func (b baoshijie)run()  {
	fmt.Printf("%s速度100迈~\n", b.brand)
}
func drive(c car)  {
	c.run()
}
func main()  {
	var f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}

