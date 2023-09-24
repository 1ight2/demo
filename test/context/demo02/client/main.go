package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println(<-stop)
				return
			default:
				fmt.Println("still working")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("stop the goroutine")
	close(stop)
	time.Sleep(5 * time.Second)
}
