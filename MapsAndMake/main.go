package main
import ("fmt")

func MakeLogic(){
  var a = make(map[string]int)
  a["First"] = 1
  a["Second"] = 2
  fmt.Println(a["First"])
  fmt.Println(a["Second"])
}

func main() {
 data := map[string]string{"Name": "John Doe","Email": "johndoe@email.com", "Password": "mypassword"}
  var data2 = map[string]int{"Cars":2,"Money":122224} 

  fmt.Println(data["Email"])
  fmt.Println(data2)
  MakeLogic()
}