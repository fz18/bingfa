package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	同步调度：
		每增加一个子协程，就向等待组中+1， 每结束一个协程就向等待组-1。主协程会阻塞等待组中协程数位0.  这种方式可以让主协程结束于
		最后一个子协程结束的时间点上。

		之前的方法用一个变量，或者使用管道。
 */

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(){
		fmt.Println("协程A开始工作")
		time.Sleep(3*time.Second)
		fmt.Println("协程Aover")
		wg.Done()
	}()
	// 计时器
	wg.Add(1)
	go func(){
		fmt.Println("协程B开始工作")
		<-time.After(5)
		fmt.Println("协程Bover")
		wg.Done()
	}()

	// ticker
	wg.Add(1)
	go func(){
		fmt.Println("协程C开始工作")
		ticker := time.NewTicker(time.Second)
		for i := 0; i < 4; i++ {
			<-ticker.C
		}
		ticker.Stop()
		fmt.Println("协程Cover")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main over")
}
