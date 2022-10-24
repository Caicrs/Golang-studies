package main
import "fmt"
func main() {
    var name = "Caic Rocha" // Type declaration is optional here.
    fmt.Printf("Variable 'name' is of type %T\n", name)
	var age, birthday, developer = 20, "18/10/2002", true
	fmt.Printf("age: %T, birthday: %T, developer: %T\n", 
	age, birthday, developer)
}