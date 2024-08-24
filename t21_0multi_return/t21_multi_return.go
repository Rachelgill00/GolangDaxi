package main

import "fmt"

func f2(ch chan int) {
	for {
		v, ok := <-ch
		//When chanel is closed, "ok" will return "false"
		if !ok {
			fmt.Println("Chanel is closed!")
			break
		}
		fmt.Printf("v:%v  ok:%v\n", v, ok)
	}
}

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}
