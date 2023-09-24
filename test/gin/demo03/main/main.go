package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

//验证器返回自定义错误
type SignUserInfo struct {
	Name string `from:"name"json:"name" binding:"required" msg:"用户名错误"`
	Age  int    `from:"age"json:"age" binding:"required,lt=30,gt=18" msg:"请输入年龄"`
}

//自定义返回错误
func GetValidMsg(err error, obj interface{}) string {
	//obj为结构体指针
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错的字段名,拿到结构体的具体字段
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回,当找到第一个错误的信息时,就可以结束
			}
		}
	}
	return err.Error()
}

func _Raw(c *gin.Context) {
	user := SignUserInfo{}
	if err := c.ShouldBindJSON(&user); err != nil {
		str := GetValidMsg(err, user)
		c.JSON(400, gin.H{"err": str, "sr": err.Error()})
	}

	c.JSON(200, gin.H{"data": &user})
}
func main() {
	Router := gin.Default()
	Router.GET("/index", _Raw)
	Router.Run(":8080")

}
