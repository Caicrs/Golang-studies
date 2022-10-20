package main
import "fmt"

func calc(a int, b int) int {

	return a + b
}

func calc2(a, b, c int) int {
	return a * b + c
}

func main() {

	res := calc(1, 2)
	fmt.Println(res)
	res = calc2(1, 2, 3)
	fmt.Println(res)
}
