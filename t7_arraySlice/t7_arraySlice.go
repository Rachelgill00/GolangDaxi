package main

import "fmt"

func printSlice(a []int) {
	fmt.Println(a)
}

func main() {
	//way1
	slice1 := []int{1, 2, 3}
	fmt.Println(slice1)

	//way2
	var slice2 []int = make([]int, 3)
	slice2[0] = 799
	fmt.Println(slice2)

	//way3
	slice4 := make([]int, 3)
	fmt.Println(slice4)

	//-----------------------------

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	printSlice(numbers)

	fmt.Println("numbers[1:4(4 is not including)] == ", numbers[1:4])

	fmt.Println("numbers[:3] == ", numbers[:3])

	fmt.Println("numbers[4:] == ", numbers[4:])

}
