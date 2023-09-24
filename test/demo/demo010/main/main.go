//实现阻塞且并发安全的map
//GO⾥⾯MAP如何实现key不存在 get操作等待 直到key存在或者超时，保证并发安全，且需要实现以下接⼝：
//type sp interface {
//	Out(key string, val interface{})                  //存入key/val，如果该key读取的 goroutine挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
//	Rd(key string, timeout time.Duration) interface{} //读取一个key，如果 key不存在阻塞，等待key存在或者超时
//}
package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mu      sync.Mutex
	condMap map[string]*sync.Cond
	data    map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		condMap: make(map[string]*sync.Cond),
		data:    make(map[string]interface{}),
	}
}

func (sm *SafeMap) Out(key string, val interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, ok := sm.condMap[key]; !ok {
		sm.condMap[key] = sync.NewCond(&sm.mu)
	}

	sm.data[key] = val
	sm.condMap[key].Broadcast()
}

func (sm *SafeMap) Rd(key string, timeout time.Duration) interface{} {
	sm.mu.Lock()
	cond, ok := sm.condMap[key]
	if !ok {
		cond = sync.NewCond(&sm.mu)
		sm.condMap[key] = cond
	}

	startTime := time.Now()
	for sm.data[key] == nil {
		remainingTime := timeout - time.Since(startTime)
		if remainingTime <= 0 {
			break
		}
		cond.WaitTimeout(remainingTime)
	}

	data := sm.data[key]
	sm.mu.Unlock()
	return data
}

func main() {
	safeMap := NewSafeMap()

	go func() {
		time.Sleep(2 * time.Second)
		safeMap.Out("key", "value")
	}()

	result := safeMap.Rd("key", 3*time.Second)
	if result != nil {
		fmt.Println(result)
	} else {
		fmt.Println("Timeout")
	}
}
