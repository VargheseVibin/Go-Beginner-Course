package main

import "fmt"

func main() {
	fmt.Println("Hello, World. Let's Go!!")

	var whatToSay string
	var i int

	whatToSay = "GoodBye, Cruel World!"
	i = 17

	fmt.Println(whatToSay)
	fmt.Println("My Date of Birth is", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
