package main
import "fmt"
// VARIABLES
var a string = "string"
var b = "This is a infered variable, string too"
var myBool bool = true
var myInt int = 2
var myFloat float32 = 2.92
// CONSTANT AND BLOCK DECLARATION
const MYCONSTANT = "Sup bro...i'am a const, no re-declarations here !!"
const (
	FIRST = "First constant"
	SECOND = 7
)


func main() {
	 a := "This is a string" // New declaration inside the function
	 myVar := "New var" // This type of declaration its different that "=" because not working when its declared outside the function
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(MYCONSTANT)
	fmt.Println(myVar)
	fmt.Println(myBool)
	fmt.Println(myInt)
	fmt.Println(myFloat)

	myInt = -2 // NEW VALUE NEGATIVE
	myFloat = -2.92 // NEW VALUE NEGATIVE
	fmt.Println(myInt)
	fmt.Println(myFloat)

	// Multiple declaration
	var first,second = "My string", true
	fmt.Println(first,second)
	third,fourth := "My string", "without variable type declaration"
	fmt.Println(third,fourth)
	fmt.Println(FIRST,SECOND)
}