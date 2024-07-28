package main

import "fmt"

func main() {
	//1 basic
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//2 equal to basic
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	//3
	for {
		fmt.Println("This will loop forever.")
		break // Without this break, it would loop indefinitely
	}

	//4
	//1)
	arr := [4]int{1, 2, 3, 4}
	for i, v := range arr {
		fmt.Printf("index : %d, value : %d \n", i, v)
	}
	//2)
	m := map[int]string{1: "a", 2: "b"}
	for k, v := range m {
		fmt.Printf("key: %v, value: %v\n", k, v)
	}

	//3)
	s := "abc"
	for i, c := range s {
		fmt.Printf("i: %v, c: %c\n", i, c)
	}

	f := "hello"
	for i, c := range f {
		fmt.Printf("Index: %d, Character: %c\n", i, c)
	}

	//5
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		if i%3 == 0 {
			continue //// Skips the rest of the loop body for even i （3不输出）
		}
		fmt.Println(i)
	}
}
