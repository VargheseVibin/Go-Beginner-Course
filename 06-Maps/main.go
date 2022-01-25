package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)

	myMap["wife"] = "Anumol"
	myMap["brother"] = "Febin"
	myMap["myIde"] = "VSCode"
	myMap["myIde"] = "PyCharm"

	log.Println(myMap["wife"])
	log.Println(myMap["brother"])
	log.Println(myMap["myIde"])

	myMap2 := make(map[string]int)
	myMap2["one"] = 1
	myMap2["two"] = 2

	log.Println(myMap2["one"])
	log.Println(myMap2["two"])

	myUserMap := make(map[string]User)
	me := User{
		FirstName: "Vibin",
		LastName:  "Varghese",
	}
	wife := User{
		FirstName: "Anu",
		LastName:  "Baby",
	}
	myUserMap["me"] = me
	myUserMap["wife"] = wife

	log.Println(myUserMap["me"].FirstName)
	log.Println(myUserMap["wife"].FirstName)

}
