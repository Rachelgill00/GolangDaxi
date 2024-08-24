package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func OG() {
	fmt.Println("One goroutine!")
	wg.Done()
}

func main() {
	wg.Add(1)
	go OG()
	wg.Wait()
	println("This is main!")
}
