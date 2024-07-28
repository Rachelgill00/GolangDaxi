package main

import "fmt"

type Person struct {
	Name string
}

// p是变量名
func (p Person) test() {
	fmt.Println("test() name=", p.Name)
}

func (p Person) count() {
	res := 0
	for i := 0; i <= 100; i++ {
		res += i
	}
	fmt.Println(p.Name, "the result of the count is:", res)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

func main() {
	var p Person
	p.Name = "tom"
	p.test()
	p.count()
	res := p.getSum(10, 20)
	fmt.Println(res)
}
