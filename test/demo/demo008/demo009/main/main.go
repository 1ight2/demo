package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	name string `json:"name"`
}

func main() {
	js := `{ 
		"name":"11" 
		}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("people: ", p)
}

//以上代码此时输出为零
//Unmarshal 将js中的json数据反序列化解析后写入了p
//people中的name属性小写为私有
//在 json 解码或转码的时候也⽆法上线私有属性的转换。
//题目中是⽆法正常得到 People 的 name 值的。⽽且，私有属性 name 也不应该加
//json 的标签
