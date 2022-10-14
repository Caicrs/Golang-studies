package main
import "fmt"

var myArray = []int{1,2,3,4,5,6,7,8,9,10,11}

func main(){
	for i:=0; i<len(myArray); i++ {
		if myArray[i] % 2 == 0 {
			fmt.Println("Its true: ", myArray[i])
		} else {fmt.Println(myArray[i])}	
	}
	
}