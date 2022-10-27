package main

import "fmt"

func main() {
	var i interface{} = "Hello World"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	g, ok := i.(bool) // panic
	fmt.Println(g, ok)
}
