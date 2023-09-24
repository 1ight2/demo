package main

import (
	"fmt"
	"reflect"
	"strings"
)

//翻转字符串
//请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字 符串(可以使用单个过程变量)。
//给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度小于等于 5000。

func reverString(s string) (string, bool) {
	if strings.Count(s, "") > 5000 {
		return s, false
	}
	str := []rune(s)
	len := len(str)
	for i := 0; i < len/2; i++ {
		str[i], str[len-1-i] = str[len-1-i], str[i]
	}
	return string(str), true
}
func reverString2(s string) (string, bool) {
	if strings.Count(s, "") > 5000 {
		return s, false
	}
	//为了能够在后续操作中修改字符。rune 切片是由字符串中的字符构成的，每个字符被表示为一个 rune
	str := []rune(s)
	len := len(str)
	for i := 0; i < len/2; i++ {
		str[i], str[len-1-i] = str[len-1-i], str[i]
	}
	return string(str), true
}
func main() {
	es, bool := reverString("asdefg")
	fmt.Println(es, bool)
	a := "ab"
	fmt.Print(reflect.TypeOf(a))
	//a[1], a[2] = a[2], a[1]
	//fmt.Println(a)

}
