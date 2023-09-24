package main

import (
	"github.com/gin-gonic/gin"
)

//验证器
type SignUserInfo struct {
	Name     string   `from:"name"json:"name" binding:"required"`
	Age      int      `from:"age"json:"age" binding:"required,lt=30,gt=18"`
	Sex      string   `from:"sex" json:"sex" binding:"required,oneof=man woman"`
	Sex2     string   `from:"sex2" json:"sex2" binding:"required,eqfield=Sex"`
	LikeList []string `json:"like_list" binding:"required,dive,startswith=like"`
	IP       string   `json:"ip" binding:"ip"`
	Url      string   `json:"url" binding:'url'`
	Uri      string   `json:"uri" binding:'uri'`
}

func Index(context *gin.Context) {
	//context.String(200, "hello")
	context.JSON(200, "hello")

}

//验证
func _Raw(c *gin.Context) {
	user := SignUserInfo{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	//c.JSON(200, user)
	c.JSON(200, gin.H{"data": user})
}
func main() {
	Router := gin.Default()
	Router.GET("/index", _Raw)
	Router.Run(":8080")

}
