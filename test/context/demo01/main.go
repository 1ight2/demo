package main

import (
	"fmt"

	"time"
)

func worker(channel chan bool, name string) {
	go func() {
		for {
			select {
			case <-channel:
				fmt.Println(name, "go the stop channel")
				return
				//两秒后返回一个时间
			case <-time.After(2 * time.Second):
				fmt.Println(name, "monitoring")

			}
		}
	}()

}

func main() {
	//context.Background()根节点
	channel := make(chan bool)

	go worker(channel, "node01")
	//go worker(channel, "node02")
	//go worker(channel, "node03")

	bool1 := time.After(2 * time.Second)
	fmt.Println(bool1)

	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	channel <- true
	time.Sleep(5 * time.Second)
}
