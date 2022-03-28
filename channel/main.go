package main

import (
	"log"

	"github.com/ryanzhangau/go-learning/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)

	// run the func after defer, as soon as the current func is done
	defer close(intChan)

	// go routin
	go CalculateValue(intChan)

	num := <-intChan

	log.Println(num)
}
