package main

import (
	"fmt"
	"time"
)

// demo2 通道误用导致的bug
func main() {
	ch := make(chan string)
	go func() {
		// 这里假设执行一些耗时的操作
		time.Sleep(3 * time.Second)
		ch <- "job result"
	}()

	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(1 * time.Second): // 较小的超时时间
		fmt.Println("Request is progressing!")
		return
	}
}
