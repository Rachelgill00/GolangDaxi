package main

import "fmt"

func Recv(c chan int) {
	r := <-c
	fmt.Printf("%d is received!\n", r)
}

func main() {
	ch := make(chan int)
	go Recv(ch)
	ch <- 10
	fmt.Println("Send successfully!")
}
