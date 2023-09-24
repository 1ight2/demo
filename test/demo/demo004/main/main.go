package main

import (
	"fmt"
	"strings"
)

//判断两个给定的字符串排序后是否⼀致
//给定两个字符串，请编写程序，确定其中⼀个字符串的字符重新排列后，能否变成另⼀
//个字符串。 这⾥规定【⼤⼩写为不同字符】，且考虑字符串重点空格。给定⼀个string
//s1和⼀个string s2，请返回⼀个bool，代表两串是否重新排列后可相同。 保证两串的
//⻓度都⼩于等于5000

//判断两个字符串中所含有的字符的数量是否一致
func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))
	if sl1 > 5000 || sl2 > 5000 || s1 != s2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) != strings.Count(s2, string(v)) {
			return false
		}
	}
	return true
}

//将字符和出现的次数存储在map中,进行对比
func areAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charCount := make(map[rune]int)

	//使用k-v结构统计个字符出现次数
	for _, r := range s1 {
		charCount[r]++
		fmt.Println(charCount)
	}
	//当出现相同字符时,对应计数减1,当为负数时表示数量对不上
	for _, r := range s2 {
		charCount[r]--
		fmt.Println(charCount)
		if charCount[r] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(areAnagrams("acdd", "aced"))
}
