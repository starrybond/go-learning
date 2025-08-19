package main

import (
	"fmt"
	"sync"
)

func main() {
	const total = 100
	const bufSize = 10 // 缓冲大小

	ch := make(chan int, bufSize)
	var wg sync.WaitGroup

	// 生产者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= total; i++ {
			ch <- i // 缓冲区满时会阻塞
		}
		close(ch) // 通知消费者没有更多数据
	}()

	// 消费者
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch { // 通道关闭后自动退出
			fmt.Println("收到：", num)
		}
	}()

	wg.Wait()
	fmt.Println("全部处理完毕")
}
