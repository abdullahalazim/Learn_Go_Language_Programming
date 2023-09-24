package main

import "fmt"

/*
	When we use ... instead of specifying the length.
	The compiler can identify the length of an array,
	based on the elements specified in the array declaration.
*/
func main() {
	fruits := [...]string{"Apple", "Banana", "Melon"} //inferred
	fmt.Println("Fruits are: ", fruits)
}
