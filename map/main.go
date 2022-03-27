package main

import "log"

type User struct {
	Firstname string
	Lastname  string
}

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	newMap := make(map[string]int)

	newMap["first"] = 1
	newMap["second"] = 2
	log.Println(newMap["first"], newMap["second"])

	userMap := make(map[string]User)
	me := User{
		Firstname: "Trevor",
		Lastname:  "Sawler",
	}

	userMap["me"] = me

	log.Println(userMap["me"].Firstname)

	// map is immutable
	// map DOES NOT maintain the order

	// create a map that CAN store ANYTHING by defining interface{}
	// use by caution 
	anotherMap := make(map[string]interface{})

	anotherMap["test"] = "anything"

}
