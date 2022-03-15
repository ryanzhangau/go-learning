package main

import "fmt"

// package level variable
var packageLevelVar string

func main() {
	var whatToSay string
	var i int

	// declare then assign value
	var newVar string
	newVar = "newVar"
	fmt.Println(newVar)

	// declare with assignment
	var newVar2 = "newVar2"
	fmt.Println(newVar2)

	// declare short-hand syntax
	// NOTE: only works within func
	newVar3 := "newVar3"
	fmt.Println(newVar3)

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
