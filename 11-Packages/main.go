package main

import (
	"fmt"

	"github.com/VargheseVibin/GoPackages/helpers"
)

func main() {
	fmt.Println("Go Packages")
	var myStruct helpers.SomeType
	myStruct.TypeNumber = 10
	fmt.Println(myStruct.TypeNumber)
}
