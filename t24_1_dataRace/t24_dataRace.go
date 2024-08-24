package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup
)

// add对全局变量x执行5000次+1操作
func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
