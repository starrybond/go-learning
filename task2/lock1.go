package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	grNum := 10      // 协程数量
	incPerGr := 1000 // 每个协程递增次数

	wg.Add(grNum)
	for i := 0; i < grNum; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incPerGr; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数器值: %d\n", counter) // 期望输出 10000
}
