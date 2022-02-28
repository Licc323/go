package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// net/http server

func f1(w http.ResponseWriter, r *http.Request)  {
	//str :=
	b, err := ioutil.ReadFile("./shimei.txt")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(b)
}

func f2(w http.ResponseWriter, r *http.Request)  {
	//对于get请求，参数都放在URL上(query param)， 请求体中是没有数据的
	queryParam := r.URL.Query() //自动帮我们识别URL中的query param
	name := queryParam.Get("name")
	age  := queryParam.Get("age")
	fmt.Println(name, age)
	fmt.Println(r.Method)
	fmt.Println(ioutil.ReadAll(r.Body))// 我在服务端打印客户端发来的请求body
	w.Write([]byte("ok"))
}
func main()  {
	http.HandleFunc("/hello/", f1)
	http.HandleFunc("/shimei/", f2)
	http.ListenAndServe("127.0.0.1:9090", nil)

}
