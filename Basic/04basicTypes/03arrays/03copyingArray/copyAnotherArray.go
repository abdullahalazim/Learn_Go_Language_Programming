package main

import "fmt"

func main() {
	country := [5]string{"Bangladesh", "India", "Japan", "Saudi Arabia", "russia"}
	fmt.Println("Main Array: ", country)
	fmt.Println("Array size", len(country))
	secondArray := country //copying existing array
	fmt.Println("After copying Array: ",secondArray)
	fmt.Println("And Size: ",len(secondArray))
}
