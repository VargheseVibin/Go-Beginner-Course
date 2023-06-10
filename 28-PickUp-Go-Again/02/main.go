package main

import "fmt"

func calcSquares(nums []int) ([]int, bool) {
	squares := []int{}
	for _, n := range nums {
		squares = append(squares, n*n)
	}
	return squares, true
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	squares, _ := calcSquares(nums)
	fmt.Println(squares)
}
