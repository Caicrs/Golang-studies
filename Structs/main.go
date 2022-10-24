package main

import "fmt"

type User struct {
	Name string
	Age int
	Email string
	Password string
	IsAdmin bool
}

func main(){
var data = User{Name:"John Doe",Age:22,Email:"johndoe@gmail.com",Password:"johndoe219",IsAdmin:true}
	fmt.Println(data)
	fmt.Println(data.Email)
	fmt.Println(data.Age)
	fmt.Println(data.IsAdmin)
}