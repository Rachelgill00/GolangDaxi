package main

import (
	"fmt"
)

func OG() {
	fmt.Println("One goroutine!")
}

func main() {
	go OG()
	println("This is main!")
}

// Output: just "This is main!"
// There;s no "One goroutine!"
// 当 main 函数结束时整个程序也就结束了，同时 main goroutine 也结束了，
// 所有由 main goroutine 创建的 goroutine 也会一同退出。
// 也就是说我们的 main 函数退出太快，
// 另外一个 goroutine 中的函数还未执行完程序就退出了，导致未打印出“One goroutine!”。
