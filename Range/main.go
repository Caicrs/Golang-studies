package main
import "fmt"
func main(){
	a := []string{"Foo", "Bar"}
	for i, s := range a {
    fmt.Println(i, s)
	// " i " its the index
	// " s " its the content inside " a " array
}
}