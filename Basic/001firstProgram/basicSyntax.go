// Copyright (c) 2023 Abdullah Al Azim
package main

import "fmt"

//single line comment
/*
multiple line comment
*/

func main() {
	mainValue := secondMain("This is me, Abdullah Al Azim")
	fmt.Println(mainValue)
}

func secondMain(msg string) string{
	return "message from main func and the message is "+msg
}