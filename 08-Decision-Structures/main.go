package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	var myNum int
	myNum = 100
	var myBool bool
	myBool = false

	if myNum >= 99 && myBool {
		log.Println("myNum is more than 99 and isTrue is True")
	}

	myAnimal := "cat1"
	// Go Breaks out of the switch after meeting the FIRST matching CASE
	switch myAnimal {
	case "cat":
		log.Println("Animal is Cat")
	case "dog":
		log.Println("Animal is Dog")
	default:
		log.Println("Animal is Unknown")
	}
}
