package main

import (
	"fmt"
	"math/rand"
	"sync"
	_ "sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			ch <- rand.Intn(100)
		}
	}()
	go func() {
		for x := range ch {
			fmt.Println(x)
			time.After()
			wg.Done()
		}
	}()
	wg.Wait()
}
