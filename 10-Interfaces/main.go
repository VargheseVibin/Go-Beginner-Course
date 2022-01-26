package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{"Fido", "GermanShepherd"}
	PrintInfo(&dog)

	cat := Cat{"Fluffy", "Asian", 26}
	PrintInfo(&cat)
}

func PrintInfo(a Animal) {
	fmt.Println("Animal says", a.Says(), " and has #", a.NumberOfLegs(), " legs")
}

func (d *Dog) Says() string {
	return "Whooff Whooff!"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (c *Cat) Says() string {
	return "Meoww Meoww!"
}

func (c *Cat) NumberOfLegs() int {
	return 4
}
