package main

import "fmt"

func findFactorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * findFactorial(n-1)
	}
}

func main() {
	f := findFactorial(4)
	fmt.Println(f)

	x := func(n int) int {
		return n * n
	}

	fmt.Printf("%T \n", x)
	fmt.Println(x(f))

	y := func(n int) int {
		return 10 * n
	}(f)

	fmt.Printf("%T \n", y)
	fmt.Println(y)
}
