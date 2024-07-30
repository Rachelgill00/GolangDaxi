package main

import "fmt"

func main() {
	var x interface{}

	var b2 float32 = 1.1

	x = b2

	y := x.(float32)
	//若执行y := x.(float64) 就会报错，不会执行以下语句
	fmt.Printf("The type if y is %T, and the value is %v\n", y, y)

	//待检测的断言：添加一个变量来判断是否成功
	if g, ok := x.(float32); ok {
		fmt.Printf("The type if g is %T, and the value is %v\n", g, g)
	}
}
