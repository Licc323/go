package main

import (
	"encoding/json"
	"fmt"
)

//json
//反射：在程序运行期间，对程序本身进行访问和修改的能力，程序在编译时变量被转换为内存地址
//变量名不会被编译写入到可执行部分，在运行时，程序无法获取自身的信息

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main()  {
	str := `{"name":"chengcheng","age":25}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
}
