package main

import (
	"fmt"

	"strings"
	"sync"
)

//交替打印数字和字母
//使⽤两个goroutine交替打印序列，⼀个goroutine打印数字,另外⼀
//个goroutine打印字⺟，最终效果如下：
//12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
//解题思路:使用channel来控制打印的进度
func main() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	//打印数字
	go func() {
		i := 1
		for true {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				break
			default:
				break
			}
		}
	}()
	wait.Add(1)
	//打印字母
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		//索引,打印到的位置
		i := 0
		for true {
			select {
			case <-letter:
				//strings.Count() 统计字符在另一个字符串中出现的次数,当substr为0,计算整个字符串中字符数
				if i >= strings.Count(str, "")-1 {
					fmt.Print(len(str))
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				//if i >= strings.Count(str, "") {
				//	i = 0
				//}
				fmt.Print(str[i : i+1])
				i++
				number <- true
				break
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()
}
