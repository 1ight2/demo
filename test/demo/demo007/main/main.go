package main

import "fmt"

type Param map[string]interface{}
type Show struct {
	Param
}

//这段代码存在一个潜在的问题，是因为在 main1 函数中，你尝试在一个空指针的 Show 结构体中的 Param 字段中添加一个键值对。
//但是在这个代码中，s 是一个指向 Show 结构体的指针，它尚未被初始化，因此不能直接操作 s.Param。
//func main() {
//	s := new(Show)
//	s.Param["RMB"] = 10000
//}
func main() {
	s := &Show{
		Param: make(Param), // 初始化 Param 字段
	}
	s.Param["RMB"] = 10000

	fmt.Println(s)
}
