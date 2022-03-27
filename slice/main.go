package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")
	log.Println(mySlice)

	// slice maintains the order
	var intSlice []int

	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 3)

	log.Println(intSlice)

	sort.Ints(intSlice)

	log.Println(intSlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	//get range from slice
	log.Println(numbers[0:2])
}
