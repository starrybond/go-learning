package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 打印奇数的协程
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数: %d\n", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}

	}()

	// 打印偶数的协程
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数: %d\n", i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
	}()

	wg.Wait()

}
