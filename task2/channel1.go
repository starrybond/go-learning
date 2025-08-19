package main

import "fmt"

func main() {
	// 创建一个无缓冲通道
	ch := make(chan int)

	// 发送方协程
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i // 发送到通道
		}
		close(ch) // 发送完毕后关闭通道
	}()

	// 接收方协程
	go func() {
		for num := range ch { // 通道关闭后自动退出循环
			fmt.Println("收到:", num)
		}
	}()

	// 等待两个 goroutine 完成（简单方式：sleep）
	// 生产代码可用 sync.WaitGroup
	select {} // 阻塞主线程，防止提前退出
}
