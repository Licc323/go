package main

import (
	_3split_string "ccLiStu.com/studgo/day09/03split_string"
	"fmt"
)

func main()  {
	ret := _3split_string.Split("babcbef", "b")
	fmt.Printf("%v\n", ret)
	ret2 := _3split_string.Split("bbb", "b")
	fmt.Printf("%v\n", ret2)

}
