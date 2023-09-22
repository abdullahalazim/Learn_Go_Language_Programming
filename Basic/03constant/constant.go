package main

import "fmt"

const Pi = 3.1415
const Size int64 = 1024
const x, y = 1, 2

func main() {
	const (
		Phi = 1.618
		E   = 2.710
	)
	const (
		Saturday = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
	)
	fmt.Println(Pi, Size, x,y, Phi,E)
	fmt.Println(Saturday,Sunday,Monday,Tuesday,Wednesday,Thursday,Friday)
}
