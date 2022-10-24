package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:5]
	fmt.Println(s)
	
	s = s[2:5]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
	
}