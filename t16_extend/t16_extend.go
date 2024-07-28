package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

// 将pupil和Graduate 共有的方法也绑定到 *Student
func (stu *Student) ShowInfo() {
	fmt.Printf("[Name: %v, Age: %v, Score: %v]\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

type Pupul struct {
	Student //原本是stu Student， 现在是匿名结构体

}

func main() {
	pupil := &Pupul{}
	//先去它的匿名结构体，再取它的字段
	pupil.Student.Name = "tom"
	pupil.Student.SetScore(98)
	pupil.Student.ShowInfo()
}
