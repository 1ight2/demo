package main

import (
	"fmt"
	"strings"
	"unicode"
)

//字符串替换问题
//请编写⼀个⽅法，将字符串中的空格全部替换为“%20”。 假定该字符串有⾜够的空间存
//放新增的字符，并且知道字符串的真实⻓度(⼩于等于1000)，同时保证字符串由【⼤⼩
//写的英⽂字⺟组成】。 给定⼀个string为原始的串，返回替换后的string。

//strings.Replace 是一个字符串处理函数，用于将字符串中的某个子串替换为另一个子串
//func Replace(s, old, new string, n int) string
//s：源字符串，需要进行替换操作的字符串。
//old：要被替换的子串。
//new：替换后的子串。
//n：指定最多替换的次数。如果 n 为正数，将最多替换 n 次；如果 n 为负数，将替换所有匹配项；如果 n 为 0，则不进行替换

func replaceBlack(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}

	for _, v := range s {
		//当非空就必须为字符
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}

func main() {
	fmt.Println(replaceBlack("a a"))
}
