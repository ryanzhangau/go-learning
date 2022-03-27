package main

import "log"

func main() {
	var isTrue bool

	isTrue = true

	if isTrue == true {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	// switch
	// Note GoLang breaks when find the matched case
	myVar := "dfdsf"
	switch myVar {
	case "cat":
		log.Println("cat")
	case "dog":
		log.Println("dog")
	case "fish":
		log.Println("fish")
	default:
		log.Println("default")
	}
}
