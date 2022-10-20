package main

import "fmt"

func calc(a int, b int) {
	total := 0
	total = a + b
	fmt.Println(total)
}

func calc2(a int, b int,c string) {
	total := 0
	total = a * b
	fmt.Println(total,c)
}

func main() {
	calc(10, 20)
	calc2(10, 20,"Cars")
}
