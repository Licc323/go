package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http client
//共用一个client适用于，请求比较频繁
var(
	client = http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: false,
		},
}
)
func main()  {
	//resp, err := http.Get("http://127.0.0.1:9090/xxx/")
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}
	data := url.Values{}
	urlObj, _ := url.Parse("http://127.0.0.1:/9090/")
	data.Set("name", "雨诗")
	data.Set("age", "20")
	queryStr := data.Encode() //URL encode之后的URL
	fmt.Println(queryStr)
	urlObj.RawQuery = queryStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	//resp, err := http.DefaultClient.Do(req)
	//if err != nil {
	//	fmt.Printf("get url failed, err:%v\n", err)
	//	return
	//}

	//请求不是特别频繁的

	//禁用KeepAlive的cilent
	//tr := &http.Transport{
	//	DisableKeepAlives: true,
	//}
	//client := http.Client{
	//	Transport: tr,
	//}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("get url failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close() //一定要记得关闭
	//从resp中把服务端返回的数据读出来
	//var data []byte
	//resp.Body.Read()
	//resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) //我在客户端读出服务端返回的响应body
	if err != nil {
		fmt.Printf("read resp.body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
