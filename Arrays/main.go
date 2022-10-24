package main

import "fmt"

var myArr = []string{"potatoes","apple"}

func main() {
	var a = []string{"RED","GREEN"}
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	
	fmt.Println(myArr)
}