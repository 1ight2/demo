package main

import (
	"fmt"
	"time"

	"math/rand"
	"sync"
)

//写代码实现两个 goroutine，其中⼀个产⽣随机数并写⼊到 go channel 中，另外⼀
//个从 channel 中读取数字并打印到标准输出。最终输出五个随机数

func main() {
	pool := make(chan int)
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	wait := sync.WaitGroup{}
	wait.Add(2)
	go func(wait *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			num := rand.Intn(5)
			pool <- num
		}
		defer close(pool)
		wait.Done()
	}(&wait)
	go func(wait *sync.WaitGroup) {
		for i := 0; i < 5; i++ {
			num := <-pool
			fmt.Println(num)
		}
		wait.Done()
	}(&wait)
	wait.Wait()
}
