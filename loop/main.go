package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// range
	animals := []string{"cat", "dog", "fish", "cow", "horse"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	// use _ to ignore the index i
	for _, animal := range animals {
		log.Println(animal)
	}

	// Loog map
	animalMap := make(map[string]string)
	animalMap["cat"] = "Fluffy"
	animalMap["dog"] = "Fido"

	for animalType, animal := range animalMap {
		log.Println(animalType, animal)
	}
}
