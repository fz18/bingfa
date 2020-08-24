package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var rwm sync.RWMutex
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			rwm.RLock()
			fmt.Println("读数据库")
			<-time.After(3*time.Second)
			rwm.RUnlock()
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			rwm.Lock()
			fmt.Println("写数据库")
			<-time.After(3*time.Second)
			rwm.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("main over")
}
