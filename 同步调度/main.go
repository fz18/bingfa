package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	需求：
	1. 30 读 10 写
	2. 并发控制5
	3. 允许一写多读
	4.定时器模拟耗时
	5.恰好在子线程结束时结束
 */

func ReadDB(wg *sync.WaitGroup, chSem chan int, rwm *sync.RWMutex){
	chSem <- 123
	rwm.RLock()
	fmt.Println("读")
	<-time.After(time.Second)
	rwm.RUnlock()
	<-chSem
	wg.Done()
}
func WriteDB(wg *sync.WaitGroup, chSem chan int, rwm *sync.RWMutex){
	chSem<-123
	rwm.Lock()
	fmt.Println("写")
	ticker := time.NewTicker(time.Second)
		<-ticker.C
	rwm.Unlock()
	<-chSem
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var rwm sync.RWMutex
	chSempher := make(chan int, 5)
	for i := 0; i < 30; i++ {
			wg.Add(1)
			go ReadDB(&wg, chSempher, &rwm)
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go WriteDB(&wg, chSempher, &rwm)
	}
	wg.Wait()
	fmt.Println("main over")
}
