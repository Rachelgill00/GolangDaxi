package main

import (
	"fmt"
	"time"
)

func OG() {
	fmt.Println("One goroutine!")
}
func main() {
	fmt.Println("This is main!")
	go OG()
	time.Sleep(time.Second)
}
