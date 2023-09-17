// Copyright (c) 2023 Abdullah Al Azim
/*
Golang Variables
bool -> used for boolean (true/false)
string -> used for character/text
int int8 int16 int32 int64 -> used for integer value
uint uint8 unit16 unit32 unit64 uintptr -> used for unsigned integer value
byte -> alias for uint8
rune -> alias for int32. represent unicode code point
float32 float64 -> used for floating point number [such as 74.445545]
complex64 complex128 -> used for complex number
*/
package main
import "fmt"

func main() {
	//Declaring variable
	var nameOfMine string = "Abdullah Al Azim"
	var age int8 = 27
	fmt.Println("Name of mine: ",nameOfMine," And age: ",age)
	//Assign multiple values
	apple,banana := 3,4
	fmt.Println("Total Apple is: ",apple," and banana is: ",banana)
	//variables names are case sensitive pickOne and PickOne is not same
	var pickOne float32 = 3.14
	var PickOne float32 = 2.366545
	fmt.Println("These two variables are: ",pickOne,PickOne)
}
