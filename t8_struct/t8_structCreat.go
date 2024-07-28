package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//way1
	var p1 Person
	fmt.Println(p1)

	//way2 {}
	p2 := Person{"mary", 12}
	fmt.Println(p2)

	//way3
	//p3 := new(Person)
	var p3 *Person = new(Person)
	(*p3).Name = "smith"
	//这样写也可以： go作者进行了优化，因为觉得上一行的写法太奇怪
	p3.Age = 13
	fmt.Println(p3)

	//way4
	//其实可以直接写
	p4 := &Person{"xiaomi", 78}
	// var p4 *Person = &Person{}
	// (*p4).Name = "xiaoming"
	// p4.Age = 77
	fmt.Println(p4)

}
