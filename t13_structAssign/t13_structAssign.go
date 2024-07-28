package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func (stu Stu) String() string {
	return fmt.Sprintf("{%s %d}", stu.Name, stu.Age)
}

func main() {
	//way1 创建时，直接写出字段值
	//1)
	var st1 = Stu{"xiaoming", 19}
	fmt.Println(st1)
	//2)
	stu2 := Stu{"xiaomei", 20}
	fmt.Println(stu2)

	//way2  创建时，写出字段名和字段值
	//这种写法，就不依赖字段的定义顺序
	//1)
	//var stu3  = Stu{      //这样写也行
	var stu3 Stu = Stu{
		Name: "Jace",
		Age:  22,
	}
	//2)
	stu4 := Stu{
		Name: "Mary",
		Age:  88,
	}
	fmt.Println(stu3, stu4)

	//way3 其实差不多，就是返回指针
	var stu5 *Stu = &Stu{
		"smith", 30,
	}
	fmt.Println(stu5)

	stu6 := &Stu{
		Name: "qin",
		Age:  23,
	}
	fmt.Println(stu6)

}
