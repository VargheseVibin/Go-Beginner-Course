package main

import "log"

func main() {
	var myColor string
	myColor = "Green"

	log.Println("myString is set to", myColor)
	changeColorUsingPointer(&myColor)
	log.Println("myString is NOW set to", myColor)
}

func changeColorUsingPointer(s *string) {
	log.Println("s is set to", s)
	newColor := "Red"
	*s = newColor
}
