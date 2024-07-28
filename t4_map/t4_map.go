package main

import "fmt"

func main() {
	//way1
	var myMap1 map[string]string
	myMap1 = make(map[string]string)

	myMap1["one"] = "java"
	myMap1["two"] = "go"

	fmt.Println(myMap1)

	//way2
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "go"

	fmt.Println(myMap2)

	myMap3 := map[string]string{
		//this mean that you have already creat the space'

		"one":   "php",
		"two":   "c++",
		"three": "python",
	}
	fmt.Println(myMap3)

	map4 := map[int]string{
		1: "java",
		2: "go",
	}

	fmt.Println(map4)

}
