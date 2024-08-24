package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //登记1个goroutine
		go func(i int) {
			defer wg.Done() //登记1个goroutine-1
			fmt.Printf("This is %d goroutine!\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("This is main")
}