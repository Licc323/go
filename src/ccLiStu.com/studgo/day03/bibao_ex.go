package main

import (
	"fmt"
	"strings"
)

// 闭包

func makeSufficFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main()  {
	jpgFunc := makeSufficFunc(".jpg")
	txtFunc := makeSufficFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
