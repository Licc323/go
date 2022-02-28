package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main()  {
	//创建路由
	//默认使用了2个中间件logger(),Recovery()
	r := gin.Default()
	//现在表单上传大小10MB

	//api参数
	r.POST("upload", func(c *gin.Context) {
		//表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//传到项目根目录，名字就用本身的
		c.SaveUploadedFile(file, file.Filename)
		//打印信息
		c.String(200, fmt.Sprintf("'%s upload!", file.Filename))
	})
	r.Run(":8080")
}
func getting(c *gin.Context)  {

}