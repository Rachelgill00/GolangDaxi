package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x       int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func writeWithLock() {
	mutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func readWithLock() {
	mutex.Lock()
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func writeWithRWLock() {
	rwMutex.Lock()
	x = x + 1
	time.Sleep(10 * time.Millisecond)
	rwMutex.Unlock()
	wg.Done()
}

func readWithRWLock() {
	rwMutex.RLock()
	time.Sleep(time.Microsecond)
	rwMutex.RUnlock()
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	// wc个并发写操作
	for i := 0; i < wc; i++ {
		wg.Add(1)
		go wf()
	}

	// rc个并发读操作
	for i := 9; i < rc; i++ {
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x : %v   cost : %v\n", x, cost)
}
func main() {
	fmt.Println("RWLock:")
	fmt.Println("Lock:")
	do(writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms
	do(writeWithLock, readWithLock, 10, 1000)     //x:10 cost:1.466500951s
}
