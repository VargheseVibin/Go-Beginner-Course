package main

import "fmt"

func main() {
	fmt.Println(addNumbers(10))
	fmt.Println(addNumbers(10, 10, 10))
	fmt.Println(addNumbers(10, 10, 10, 10, 10))
}

func addNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum = sum + value
	}
	return sum
}
