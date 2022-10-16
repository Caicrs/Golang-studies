package main

import (
	"errors"
	"fmt"
)

var result = 1

func chain(n int) {
	defer func() {
		if r := recover(); r != nil {  // recover() = make a re-catch for all panic functions, this is good for no rewrite all the errors log
			fmt.Println(r," | defer func")
		}
	}()

	if n == 0 {
		panic(errors.New("Cannot multiply a number by zero")) // panic(errors.New(" text here !")) = Serves to create a error log and signalize for the user
	}else if n < 0 {
		panic(errors.New("Its minor than 0"))
	}else {
		result *= n
		fmt.Println("Output: ", result)
	}
}

func main() {
	chain(5)
	chain(0)
	chain(-2)
	chain(8)
}