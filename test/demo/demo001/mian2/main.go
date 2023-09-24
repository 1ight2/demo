package main

import (
	"fmt"
	"sync"
)

func main() {
	letter, number := make(chan bool), make(chan bool)
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	wait := sync.WaitGroup{}
	wait.Add(2)
	go func(wait *sync.WaitGroup) {
		i := 1
		for true {
			select {
			case <-number:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letter <- true
				if i > 28 {
					wait.Done()
					return
				}
			default:
				break
			}
		}
	}(&wait)

	go func(wait *sync.WaitGroup) {
		i := 0
		for true {
			select {
			case <-letter:
				if i > len(str)-1 {
					wait.Done()
					return
				}
				fmt.Print(str[i : i+1])
				i++
				fmt.Print(str[i : i+1])
				i++
				number <- true
			default:
				break
			}
		}
	}(&wait)
	number <- true
	wait.Wait()

}
