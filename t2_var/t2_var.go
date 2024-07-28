package main

import "fmt"

// global variable
var gA int = 100
var gB int = 120

func main() {
	//way1
	var a int = 100
	fmt.Printf("a = %v\n", a)
	//way2
	var b int
	fmt.Printf("b = %v\n", b)

	//way3
	var c = 300
	fmt.Printf("c = %v\n", c)

	//way4
	d := 400
	fmt.Printf("d = %v\n", d)

	fmt.Printf("gA = %v\n", gA)
	fmt.Printf("gB = %v\n", gB)

}
