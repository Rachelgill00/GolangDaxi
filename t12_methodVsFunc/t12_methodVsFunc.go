package main

import "fmt"

type Person struct {
	Name string
	//Age  int
}

// Function
func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func main() {
	p := Person{"tom"}
	test01(p)
	//test01(&p) //error
	test02(&p)
	//test02(p) //error
}
