package main

import (
	"log"

	"github.com/VargheseVibin/GoPackages02/helpers"
)

const ceilVal = 100

func CalculateValue(intChan chan int) {
	rNum := helpers.GetRandomNumber(ceilVal)
	intChan <- rNum
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)
	myRandomNumber := <-intChan

	log.Println(myRandomNumber)
}
