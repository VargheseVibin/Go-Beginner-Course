package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Vibin")
	mySlice = append(mySlice, "Anu")

	log.Println(mySlice)

	var myIntSlice []int
	myIntSlice = append(myIntSlice, 4)
	myIntSlice = append(myIntSlice, 1)
	myIntSlice = append(myIntSlice, 2)

	sort.Ints(myIntSlice)

	log.Println(myIntSlice)

	myNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(myNumbers)
	log.Println(myNumbers[6:9])

}
