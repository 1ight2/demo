package main

import "fmt"

// switch type的使用
//在这个示例中，msg := v.(type) 用于将 v 转换为具体的类型，然后根据具体的类型执行不同的操作。
//在不同的 case 分支中，你可以使用 msg 变量来操作接口值的底层类型。
//需要注意的是，msg := v.(type) 必须在 switch 语句中使用，并且只能在类型选择中使用。
//它不能用于常规的类型断言，例如 x := v.(int)。这是因为类型选择是用于在多个类型之间选择分支，而不是用于确定一个固定的类型。

type student struct {
	Name string
}

//func zhoujielun(v interface{}) {
//	switch msg := v.(type) {
//	case *student, student:
//		msg.Name
//	}
//}

//完善后
func zhoujielun(v interface{}) string {
	switch msg := v.(type) {
	case *student:
		fmt.Println(msg)
		return msg.Name // 返回 student 类型的 Name 字段
	case student:
		fmt.Println(msg)
		return msg.Name // 返回 student 类型的 Name 字段
	default:
		return "Unknown"
	}

}

func main() {
	s := student{Name: "Zhou Jielun"}
	result := zhoujielun(s)
	println(result) // 输出 "Zhou Jielun"
}
