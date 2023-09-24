package main

import "fmt"

func main() {
	var multiArray = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	var i,j int
	for i=0; i<3; i++{
		for j=0; j<2; j++{
			fmt.Printf("a[%d][%d]: %d\n",i,j,multiArray[i][j])
		}
	}
}
