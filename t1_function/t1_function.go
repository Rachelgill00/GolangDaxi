package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// return multi value,anonymous
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 666, 777
}

func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a, "b = ", b)
	//before assignment,r1,r2=?????
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	//return result is r1=0,r2=0
	//it indicates that before assignment r1,r2 are as the Local variable
	//and a,b are foo3's parameter, r1,r2 are actually foo3's parameter too
	//and r1,r2's values defualtly initialize to 0
	//assignment to the parameters which have a name to return
	r1 = 100
	r2 = 2000
	return
	//and we also can return like bellow:
	//return 100,2000

}

func main() {
	//foo1
	c := foo1("abc", 555)

	fmt.Println("c = ", c)
	fmt.Println()

	//foo2
	ret1, ret2 := foo2("hahaha", 999)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

	//foo3
	//modify the value of ret1, ret2
	ret1, ret2 = foo3("foo3", 333)
	fmt.Println("ret1 = ", ret1, "ret2 = ", ret2)

}
