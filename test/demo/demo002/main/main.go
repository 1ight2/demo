package main

import (
	"fmt"
	"strings"
)

//判断字符串中字符是否全都不同
//请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允许使⽤额外的存储结构】。
//给定⼀个string，请返回⼀个bool值,true代表所有字符全都不同，false代表存在相同的字符。
//保证字符串中的字符为【ASCII字符】。字符串的⻓度⼩于等于【3000】。
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	i := 0
	for true {
		if i > len(s)-1 {
			return true
		}
		if strings.Count(s, s[i:i+1]) > 1 {
			return false
		}
		i++
	}
	return true
}
func isUniqueString2(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		//ASCII表一共127位
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}
func isUniqueString3(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for k, v := range s {
		//ASCII表一共127位
		if v > 127 {
			return false
		}
		//返回第一次找到索引的位置,找不到返回-1
		//当存在重复,第二个值的k一定不等于strings.Index(s, string(k))
		if strings.Index(s, string(k)) != k {
			return false
		}
		//同上,当存在重复,第一个值的k一定不等于strings.Index(s, string(k))
		//if strings.LastIndex(s, string(k)) != k {
		//	return false
		//}

	}
	return true
}
func test(s string) {
	for k, v := range s {
		fmt.Println(k, string(v), v)

	}
}
func main() {
	fmt.Print(isUniqueString2("ad."))
	test("dsada")
}
