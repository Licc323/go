package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	//1.创建路由
	r := gin.Default()
	//r.SetTrustedProxies([]string{"192.168.1.2"})
	//2.绑定路由规则，执行的函数
	//gin.Context,封装了request和response
	r.POST("/from", func(c *gin.Context) {
		//name := c.Param("name")
		//action := c.Param("action")
		//c.String(200, name+" is "+action)
		//fmt.Printf("ClinetIP: %s\n", c.ClientIP())

		name := c.DefaultQuery("name", "shimei")
		c.String(200, fmt.Sprintf("Hello %s", name))
	})
	r.POST("/xxxPost",getting)
	r.PUT("/xxxPut")
	//监听接口， 默认在8080
	r.Run(":8080")
}
//url参数可以通过defaultQuery方法获取
//
func getting(c *gin.Context)  {

}
//