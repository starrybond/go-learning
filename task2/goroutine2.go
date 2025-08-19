package main

import (
	"fmt"
	"sync"
	"time"
)

// 任务元数据：任务 + 名字（便于日志）
type task struct {
	name string
	fn   func()
}

// TaskScheduler 简易并发调度器
type TaskScheduler struct{}

// Run 并发执行所有任务并打印耗时
func (ts *TaskScheduler) Run(tasks []task) {
	var wg sync.WaitGroup
	wg.Add(len(tasks))

	for _, t := range tasks {
		// 在 goroutine 中执行
		go func(t task) {
			defer wg.Done()
			start := time.Now()
			t.fn() // 真正执行任务
			fmt.Printf("任务 %q 完成，耗时 %v\n", t.name, time.Since(start))
		}(t)
	}
	wg.Wait()
}

func main() {
	// === 构造示例任务 ===
	tasks := []task{
		{"job-A", func() { time.Sleep(300 * time.Millisecond) }},
		{"job-B", func() { time.Sleep(500 * time.Millisecond) }},
		{"job-C", func() { time.Sleep(200 * time.Millisecond) }},
	}

	// === 调度 ===
	scheduler := &TaskScheduler{}
	scheduler.Run(tasks)
}
