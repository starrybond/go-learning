package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 // 被原子操作的共享变量
	const (
		grNum    = 10   // 协程数量
		loopPerG = 1000 // 每个协程递增次数
	)

	var wg sync.WaitGroup
	wg.Add(grNum)

	for i := 0; i < grNum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < loopPerG; j++ {
				atomic.AddInt64(&counter, 1) // 无锁递增
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", atomic.LoadInt64(&counter)) // 期望 10000
}
