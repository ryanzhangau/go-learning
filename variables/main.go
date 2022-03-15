package main

import "fmt"

// package level variable
var packageLevelVar string

func main() {
	var whatToSay string
	var i int

	whatToSay = "Gooodbye, cruel world"
	fmt.Println(whatToSay)

	i = 8
	fmt.Println("i is set to", i)

	whatWasSaid := saySomething()
	fmt.Println(whatWasSaid)
}

func saySomething() string {
	return "something"
}
