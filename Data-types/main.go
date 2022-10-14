package main
import "fmt"

// boolean | true or false
var a bool = true
// string 
var b string = "its a string"


/* NUMERALS INTEGER
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
*/
// numerals | integers ( positive and negative )
var c int = 7
// numerals | int8 or byte ( the same thing ) = This integer value is of 8 bits and it represents one byte i.e number between 0-255).
var d byte = 255
var e int8 = 127


// numerals | Float32 
var f float32 = 12.2411241
var g float32 = -12.2411241

// numerals | Float64 Have the twice size of float32 but can show more high and precise numbers
var h float64 = 12.24112413525252352352235252
var i float64 = -12.241124125252523525252352362


func main(){
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
}