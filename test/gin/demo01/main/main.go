package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	//context.String(200, "hello")
	context.JSON(200, "hello")

}

//原始参数
func _Raw(c *gin.Context) {
	str := c.GetHeader("User-Agent")
	str2 := c.Request
	//str3, _ := json.Marshal(str2)
	fmt.Println(string(str))
	c.Header("token", "token")

	c.JSON(200, str2.Header)
}
func main() {
	Router := gin.Default()
	Router.GET("/index", _Raw)
	Router.Run(":8080")

}
