package main

import (
	"fmt"
	"reflect"
)

func MyAdd(a int,b int)int{
	return a + b
}
func MyAdd2(a int,b int)int{
	return a - b
}
func CallAdd(f func(a int,b int) int)  {
	v:= reflect.ValueOf(f)
	//判断v的类型是不是func
	if v.Kind() != reflect.Func{
		return
	}
	argv := make([]reflect.Value,2)
	argv[0] = reflect.ValueOf(1)
	argv[1] = reflect.ValueOf(2)
	result := v.Call(argv)
	fmt.Println(result[0].Int())
}
func main() {
	CallAdd(MyAdd)
	CallAdd(MyAdd2)
	ch := make(chan string)
}