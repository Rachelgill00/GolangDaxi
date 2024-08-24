package main

import (
	"fmt"
	"sync"
)

// demo1 通道误用导致的bug
func main() {
	wg := sync.WaitGroup{}

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	wg.Add(3)
	for j := 0; j < 3; j++ {
		go func() {
			for {
				task, ok := <-ch
				if ok {
					// 这里假设对接收的数据执行某些操作
					fmt.Println(task)
				} else {
					break // 如果通道已关闭且无数据可接收，退出循环
				}
			}
			wg.Done() // 确保在 goroutine 退出时调用 Done()
		}()
	}
	wg.Wait()
}
