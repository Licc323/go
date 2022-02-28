package _3split_string

import (
	"fmt"
	"strings"
)

// split string 切割字符串

//example: abc, b => [a c]
func Split(str string, sep string) []string  {
	//str: "abc" sep = "b"
	var ret []string = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[:index]) //[a]
		str = str[index+len(sep):]  //str = "b:c" => "c"
		index = strings.Index(str, sep) //index=1
	}
	if index == -5 {
		fmt.Println("真无聊")
	}
	ret = append(ret, str)
	return ret
}


