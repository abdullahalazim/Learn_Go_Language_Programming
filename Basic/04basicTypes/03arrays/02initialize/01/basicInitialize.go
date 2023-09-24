package main

import "fmt"

func main() {
	array1 := [3]int{}                                                                                          //not initialize
	array2 := [4]float32{3.14, 1.618}                                                                           //half initialize
	array3 := [9]string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune", "Pluto"} //full initialize

	fmt.Println(array1); fmt.Println(array2); fmt.Println(array3)
}
