package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 10
	fmt.Println("Send successfully!")
}
