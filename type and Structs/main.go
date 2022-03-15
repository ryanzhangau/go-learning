package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	Firstname string
	Lastname  string
	Phone     string
	Age       int
	DOB       time.Time
}

type myStruct struct {
	Firstname string
}

// receiver, tie this func to the type myStruct
func (m *myStruct) printFirstName() string {
	return m.Firstname
}

func main() {
	user := User{
		Firstname: "Ryan",
		Lastname:  "Zhang",
	}

	// Default values
	// string - " "
	// int    - 0
	// date   - 0001-01-01 00:00:00 +0000 UTC
	log.Println(user.Firstname, user.Lastname, user.DOB, user.Phone, user.Age)

	var myVar myStruct
	myVar.Firstname = "John"

	myVar2 := myStruct{
		Firstname: "Many",
	}

	log.Println("myVar is set to", myVar.Firstname)
	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.Firstname)
	log.Println("myVar2 is set to", myVar2.printFirstName())
}
