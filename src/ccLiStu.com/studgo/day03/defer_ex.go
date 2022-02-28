package main

import "fmt"

func calc2(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main()  {
	a := 1
	b := 2
	defer calc2("1", a, calc2("10", a, b))
	a = 0
	defer calc2("2", a, calc2("20", a, b))
	b = 1
}
