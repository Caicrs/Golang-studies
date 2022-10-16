package main
import "fmt"

func calculator1(x,y float32,command string){
	if command == "*"{
		fmt.Println(x * y)
	} else if command == "/"{
		fmt.Println(x / y)
	} else if command == "+"{
		fmt.Println(x + y)
	} else if command == "-"{
		fmt.Println(x - y)
	}else{
		fmt.Println("Error : check your request again")
	}
}

func userRole(role string){
	if role == "admin" {
		fmt.Println("Its Admin")
		} else{fmt.Println("Not is an admin")} 
}

func userRequest(request string){
	switch request {
	case "POST":
			fmt.Println("Its a POST request")
	case "GET":
			fmt.Println("Its a GET request")
	case "GETBYID":
			fmt.Println("Its a GETBYID request")
	case "DELETE":
			fmt.Println("Its a DELETE request")
	default:
			fmt.Println("Error : check your useRequest again ")	
}}

func main(){
	calculator1(2,45,"-3")
	userRole("admin")
	userRequest("GET")
}

