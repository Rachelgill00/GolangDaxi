package main

import "fmt"

type Circle struct {
	radius float64
}

func (c *Circle) area2() float64 {
	fmt.Println("area2 c address:", &c)
	return 3.14 * (*c).radius * (*c).radius

}

func main() {
	var c Circle
	fmt.Printf("main c address:= %p", &c)
	fmt.Println("main c address:", &c)

	c.radius = 5.0

	res2 := (&c).area2()
	fmt.Println(res2)

}
