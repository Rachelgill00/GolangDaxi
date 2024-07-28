package main

import (
	"fmt"

	"github.com/Rachelgill00/GolangDaxi/t14factory/model"
)

func main() {
	// var stu = model.Student{
	// 	Name:  "tom",
	// 	Score: 78.9,
	// }

	//当结构体是小写
	var stu = model.NewStudent("tom", 88.8)
	fmt.Println(stu)
}
