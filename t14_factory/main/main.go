package main

import (
	"GolangDaxi/t14_factory/model"
	"fmt"
)

func main() {
	// var stu = student{
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }
	var stu = model.NewStudent("tom~", 88.8)

	fmt.Println(stu)
}
